package main

import "fmt"

// 704. Binary Search
// Time: O(log N), Space: O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Classic Binary Search
- **Left Pointer**: Start of search range
- **Right Pointer**: End of search range
- **Middle Element**: Compare with target to narrow search
- **Range Reduction**: Halve search space each iteration

## 2. PROBLEM CHARACTERISTICS
- **Sorted Array**: Input array is sorted in ascending order
- **Unique Elements**: No duplicate elements (though algorithm works with duplicates)
- **Index Search**: Find index of target element
- **Logarithmic Time**: O(log N) time complexity

## 3. SIMILAR PROBLEMS
- Search in Rotated Sorted Array (LeetCode 33)
- Find First and Last Position of Element (LeetCode 34)
- Search Insert Position (LeetCode 35)
- Sqrt(x) (LeetCode 69)

## 4. KEY OBSERVATIONS
- **Monotonic property**: Sorted array enables binary search
- **Divide and conquer**: Halve search space each iteration
- **Boundary conditions**: Handle empty array and single element
- **Integer overflow**: Use left + (right-left)/2 instead of (left+right)/2

## 5. VARIATIONS & EXTENSIONS
- **Search Insert Position**: Find where to insert target
- **Lower/Upper Bound**: Find first/last occurrence
- **Rotated Array**: Handle rotated sorted arrays
- **Multiple Targets**: Find all occurrences

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is array sorted? Are there duplicates?"
- Edge cases: empty array, single element, target not found
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(log N) - halving search space

## 7. COMMON MISTAKES
- Using (left+right)/2 causing integer overflow
- Wrong loop condition (should be <= not <)
- Off-by-one errors in pointer updates
- Not handling empty array properly
- Forgetting to return -1 when not found

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(log N) time, O(1) space
- **Recursive version**: Can be implemented recursively but uses stack space
- **Template method**: Use consistent binary search template
- **Early termination**: Not applicable (need to verify absence)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a name in a phone book:**
- You have a phone book sorted alphabetically (sorted array)
- You want to find a specific person's name (target)
- Instead of reading from start to finish, open to the middle
- If the name comes before your target, search the first half
- If the name comes after your target, search the second half
- Repeat this process until you find the name or run out of pages

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array of integers and target value
2. **Goal**: Find index of target in array
3. **Output**: Index if found, -1 if not found
4. **Constraint**: Array is sorted in ascending order

#### Phase 2: Key Insight Recognition
- **"Divide and conquer"** → Can eliminate half the elements each step
- **"Monotonic property"** → Sorted array enables efficient search
- **"Middle comparison"** → Compare middle element to target
- **"Range reduction"** → Narrow search space based on comparison

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use two pointers, left and right, to define my search range.
I'll look at the middle element and compare it to my target.
If the middle element is my target, I'm done!
If it's smaller, my target must be in the right half.
If it's larger, my target must be in the left half.
I'll repeat this until I find it or the range is empty."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 (no elements to search)
- **Single element**: Check if it equals target
- **Target not found**: Return -1 when left > right
- **Integer overflow**: Use safe middle calculation

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[-1,0,3,5,9,12], target=9

Human thinking:
"I'll search for 9 in the sorted array:

Initial range: left=0, right=5 (entire array)
Middle index: (0+5)/2 = 2, nums[2]=3
3 < 9, so target must be in right half
New range: left=3, right=5

Middle index: (3+5)/2 = 4, nums[4]=9
9 == 9, found it! Return index 4"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Sorted array allows elimination of half elements
- **Why O(log N) time**: Each iteration halves the search space
- **Why O(1) space**: Only need pointers, no extra data structures
- **Why safe middle calculation**: Prevents integer overflow

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just linear search?"** → Binary search is much faster for large arrays
2. **"Should I use recursion?"** → Iterative is usually better (no stack overflow)
3. **"What about duplicates?"** → Algorithm still works, returns one occurrence
4. **"Can I use (left+right)/2?"** → Risk of overflow, use left + (right-left)/2

### Real-World Analogy
**Like finding a page number in a book:**
- You have a book with numbered pages (sorted array)
- You want to find a specific page number (target)
- Open to the middle of the book
- If your page is higher, look in the second half
- If your page is lower, look in the first half
- Keep dividing until you find your page

### Human-Readable Pseudocode
```
function binarySearch(array, target):
    left = 0
    right = length(array) - 1
    
    while left <= right:
        middle = left + (right - left) // 2
        
        if array[middle] == target:
            return middle
        else if array[middle] < target:
            left = middle + 1
        else:
            right = middle - 1
    
    return -1  // Target not found
```

### Execution Visualization

### Example: nums=[-1,0,3,5,9,12], target=9
```
Array: [-1][0][3][5][9][12]
Index:   0  1  2  3  4  5

Initial: left=0, right=5
mid = 0 + (5-0)/2 = 2, nums[2]=3
3 < 9 → search right half
left=3, right=5

Step 2: left=3, right=5
mid = 3 + (5-3)/2 = 4, nums[4]=9
9 == 9 → FOUND! Return 4
```

### Example: Target not found: nums=[-1,0,3,5,9,12], target=2
```
Initial: left=0, right=5
mid=2, nums[2]=3 > 2 → search left half
left=0, right=1

Step 2: left=0, right=1
mid=0, nums[0]=-1 < 2 → search right half
left=1, right=1

Step 3: left=1, right=1
mid=1, nums[1]=0 < 2 → search right half
left=2, right=1

left > right → NOT FOUND, return -1
```

### Key Visualization Points:
- **Range boundaries**: left and right pointers define search space
- **Middle calculation**: Safe formula prevents overflow
- **Range reduction**: Eliminate half the elements each iteration
- **Termination**: Stop when found or range is empty

### Memory Layout Visualization:
```
Array: [-1][0][3][5][9][12]
Index:   0  1  2  3  4  5
        ^           ^
     left=0      right=5
        mid=2 (value=3)
```

### Time Complexity Breakdown:
- **Each iteration**: O(1) - constant work
- **Number of iterations**: O(log N) - halving search space
- **Total time**: O(log N)
- **Space**: O(1) - only pointers and variables

### Alternative Approaches:

#### 1. Linear Search (O(N) time, O(1) space)
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
- **Pros**: Simple, works for unsorted arrays
- **Cons**: O(N) time, much slower for large arrays

#### 2. Recursive Binary Search (O(log N) time, O(log N) space)
```go
func search(nums []int, target int) int {
    return binarySearchHelper(nums, target, 0, len(nums)-1)
}

func binarySearchHelper(nums []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        return binarySearchHelper(nums, target, mid+1, right)
    } else {
        return binarySearchHelper(nums, target, left, mid-1)
    }
}
```
- **Pros**: Elegant, same time complexity
- **Cons**: Uses O(log N) stack space, risk of stack overflow

#### 3. Standard Library (O(log N) time, O(1) space)
```go
import "sort"

func search(nums []int, target int) int {
    idx := sort.SearchInts(nums, target)
    if idx < len(nums) && nums[idx] == target {
        return idx
    }
    return -1
}
```
- **Pros**: Uses standard library, tested
- **Cons**: Less learning value for interviews

### Extensions for Interviews:
- **Search Insert Position**: Find where to insert target
- **Lower Bound**: Find first element >= target
- **Upper Bound**: Find first element > target
- **Multiple Occurrences**: Find range of all target positions
*/
func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9},
		{[]int{-1, 0, 3, 5, 9, 12}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{}, 1},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{-10, -5, 0, 5, 10}, -5},
		{[]int{-10, -5, 0, 5, 10}, 0},
		{[]int{2, 4, 6, 8, 10}, 7},
	}
	
	for i, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Index: %d\n", 
			i+1, tc.nums, tc.target, result)
	}
}
