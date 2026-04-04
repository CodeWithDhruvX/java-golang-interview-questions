package main

import "fmt"

// 34. Find First and Last Position of Element in Sorted Array
// Time: O(log N), Space: O(1)
func searchRange(nums []int, target int) []int {
	return []int{findFirstOccurrence(nums, target), findLastOccurrence(nums, target)}
}

func findFirstOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findLastOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search for Boundaries
- **Lower Bound**: Find first occurrence using modified binary search
- **Upper Bound**: Find last occurrence using modified binary search
- **Boundary Adjustment**: Continue searching after finding target
- **Range Result**: Return [first, last] or [-1, -1] if not found

## 2. PROBLEM CHARACTERISTICS
- **Sorted Array**: Input array is sorted in ascending order
- **Duplicates Allowed**: Array may contain duplicate elements
- **Range Search**: Need both first and last positions
- **Logarithmic Time**: O(log N) solution required

## 3. SIMILAR PROBLEMS
- Search Insert Position (LeetCode 35)
- Find Peak Element (LeetCode 162)
- Count Complete Tree Nodes (LeetCode 222)
- Kth Smallest Element in a Sorted Matrix (LeetCode 378)

## 4. KEY OBSERVATIONS
- **Two separate searches**: Need first and last positions
- **Modified binary search**: Continue after finding target
- **Boundary conditions**: Handle edge cases for first/last elements
- **Early termination**: Can return early if target not found

## 5. VARIATIONS & EXTENSIONS
- **Count Occurrences**: Return count instead of range
- **Range Query**: Find elements within a value range
- **Multiple Targets**: Find ranges for multiple values
- **Insert Position**: Find where to insert new element

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are there duplicates? What if target not found?"
- Edge cases: empty array, single element, no duplicates, all same
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(log N) - two binary searches

## 7. COMMON MISTAKES
- Using standard binary search without boundary adjustment
- Not handling case where target doesn't exist
- Off-by-one errors in pointer updates
- Confusing first and last search logic
- Not checking both bounds separately

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(log N) time, O(1) space
- **Combined search**: Can find both bounds in one pass (more complex)
- **Template method**: Use consistent boundary search pattern
- **Early exit**: Return [-1, -1] if first search fails

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the first and last occurrence of a word in a dictionary:**
- You have a dictionary with words in alphabetical order
- Some words appear multiple times (duplicates)
- You want to find the first page where the word appears
- And also the last page where the word appears
- Use binary search to find any occurrence, then search boundaries
- For first occurrence, keep searching left even after finding the word
- For last occurrence, keep searching right even after finding the word

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array with possible duplicates and target value
2. **Goal**: Find first and last positions of target
3. **Output**: [first_index, last_index] or [-1, -1] if not found
4. **Constraint**: Must be O(log N) time complexity

#### Phase 2: Key Insight Recognition
- **"Two separate problems"** → Need first and last positions separately
- **"Modified binary search"** → Continue searching after finding target
- **"Boundary logic"** → Different logic for first vs last
- **"Early optimization"** → Can stop if first search fails

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find both the first and last occurrence of the target.
I'll use two separate binary searches:
1. For the first occurrence, when I find the target, I'll continue searching left
2. For the last occurrence, when I find the target, I'll continue searching right
This way I can find the exact boundaries of the target's range."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return [-1, -1] (no elements)
- **Single element**: Check if it equals target
- **No duplicates**: First and last will be the same
- **Target not found**: Return [-1, -1]

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[5,7,7,8,8,10], target=8

Human thinking:
"I'll find both boundaries for 8:

First Occurrence:
Initial: left=0, right=5
mid=2, nums[mid]=7 < 8 → search right half
left=3, right=5

mid=4, nums[mid]=8 == target → record result, search left half
right=3

mid=3, nums[mid]=8 == target → record result, search left half
right=2

left > right → First occurrence is at index 3

Last Occurrence:
Initial: left=0, right=5
mid=2, nums[mid]=7 < 8 → search right half
left=3, right=5

mid=4, nums[mid]=8 == target → record result, search right half
left=5

mid=5, nums[mid]=10 > 8 → search left half
right=4

left > right → Last occurrence is at index 4

Final result: [3, 4]"
```

#### Phase 6: Intuition Validation
- **Why two searches needed**: First and last require different boundary logic
- **Why O(log N) each**: Each is a modified binary search
- **Why continue after finding**: Need to find exact boundaries
- **Why record results**: Need to remember found positions

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just find one and expand?"** → Could be O(N) in worst case
2. **"Should I use the same logic for both?"** → No, first and last need different directions
3. **"What if target not found?"** → Both searches will return -1
4. **"Can I combine both searches?"** → Possible but more complex

### Real-World Analogy
**Like finding the first and last house number on a street:**
- You have houses numbered in order (sorted array)
- Some numbers might be repeated (duplicates)
- You want the first house with your target number
- And the last house with that same number
- Use binary search to find any house with that number
- Then search left to find the first one
- And search right to find the last one

### Human-Readable Pseudocode
```
function findRange(array, target):
    first = findFirstOccurrence(array, target)
    if first == -1:
        return [-1, -1]  // Target not found
    last = findLastOccurrence(array, target)
    return [first, last]

function findFirstOccurrence(array, target):
    left = 0
    right = length(array) - 1
    result = -1
    
    while left <= right:
        middle = left + (right - left) // 2
        
        if array[middle] == target:
            result = middle
            right = middle - 1  # Continue searching left
        else if array[middle] < target:
            left = middle + 1
        else:
            right = middle - 1
    
    return result

function findLastOccurrence(array, target):
    left = 0
    right = length(array) - 1
    result = -1
    
    while left <= right:
        middle = left + (right - left) // 2
        
        if array[middle] == target:
            result = middle
            left = middle + 1   # Continue searching right
        else if array[middle] < target:
            left = middle + 1
        else:
            right = middle - 1
    
    return result
```

### Execution Visualization

### Example: nums=[5,7,7,8,8,10], target=8
```
Array: [5][7][7][8][8][10]
Index:   0  1  2  3  4  5

=== First Occurrence ===
Initial: left=0, right=5
mid=2, nums[mid]=7 < 8 → search right half
left=3, right=5

mid=4, nums[mid]=8 == target → result=4, search left half
right=3

mid=3, nums[mid]=8 == target → result=3, search left half
right=2

left > right → First = 3

=== Last Occurrence ===
Initial: left=0, right=5
mid=2, nums[mid]=7 < 8 → search right half
left=3, right=5

mid=4, nums[mid]=8 == target → result=4, search right half
left=5

mid=5, nums[mid]=10 > 8 → search left half
right=4

left > right → Last = 4

Result: [3, 4]
```

### Example: Target not found: nums=[5,7,7,8,8,10], target=6
```
=== First Occurrence ===
Initial: left=0, right=5
mid=2, nums[mid]=7 > 6 → search left half
right=1

mid=0, nums[mid]=5 < 6 → search right half
left=1, right=1

mid=1, nums[mid]=7 > 6 → search left half
right=0

left > right → First = -1

Since first is -1, return [-1, -1]
```

### Key Visualization Points:
- **Boundary searches**: Two separate binary searches
- **Direction logic**: First searches left, last searches right
- **Result tracking**: Record found positions during search
- **Early termination**: Can stop if first search fails

### Memory Layout Visualization:
```
Array: [5][7][7][8][8][10]
Index:   0  1  2  3  4  5
                 ^  ^
               first last
               (3) (4)
```

### Time Complexity Breakdown:
- **First search**: O(log N) - modified binary search
- **Last search**: O(log N) - modified binary search
- **Total time**: O(log N) - constant factor of 2
- **Space**: O(1) - only pointers and variables

### Alternative Approaches:

#### 1. Single Pass Combined Search (O(log N) time, O(1) space)
```go
func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    first, last := -1, -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            // Found target, expand to find boundaries
            first, last = mid, mid
            // Expand left
            for first > 0 && nums[first-1] == target {
                first--
            }
            // Expand right
            for last < len(nums)-1 && nums[last+1] == target {
                last++
            }
            return []int{first, last}
        }
    }
    
    return []int{-1, -1}
}
```
- **Pros**: Single pass through array
- **Cons**: Can be O(N) in worst case if many duplicates

#### 2. Linear Scan (O(N) time, O(1) space)
```go
func searchRange(nums []int, target int) []int {
    first, last := -1, -1
    for i, num := range nums {
        if num == target {
            if first == -1 {
                first = i
            }
            last = i
        }
    }
    return []int{first, last}
}
```
- **Pros**: Simple, works for any array
- **Cons**: O(N) time, doesn't leverage sorted property

#### 3. Standard Library (O(log N) time, O(1) space))
```go
import "sort"

func searchRange(nums []int, target int) []int {
    first := sort.SearchInts(nums, target)
    if first == len(nums) || nums[first] != target {
        return []int{-1, -1}
    }
    
    last := sort.SearchInts(nums, target+1) - 1
    return []int{first, last}
}
```
- **Pros**: Uses standard library, concise
- **Cons**: Less learning value for interviews

### Extensions for Interviews:
- **Count Occurrences**: Return last - first + 1
- **Range Query**: Find all elements in [low, high] range
- **Insert Position**: Find where to insert new element
- **Multiple Targets**: Find ranges for multiple values simultaneously
*/
func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{-3, -2, -1, 0, 1}, -1},
		{[]int{1, 3, 5, 7, 9}, 4},
	}
	
	for i, tc := range testCases {
		result := searchRange(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Range: %v\n", 
			i+1, tc.nums, tc.target, result)
	}
}
