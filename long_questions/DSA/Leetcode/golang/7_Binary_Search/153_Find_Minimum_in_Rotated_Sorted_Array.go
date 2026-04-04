package main

import "fmt"

// 153. Find Minimum in Rotated Sorted Array
// Time: O(log N), Space: O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	
	// If array is not rotated
	if nums[left] <= nums[right] {
		return nums[left]
	}
	
	for left < right {
		mid := left + (right-left)/2
		
		// Check if mid is the minimum element
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		
		// Check if mid+1 is the minimum element
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		
		// Decide which half to search
		if nums[mid] >= nums[left] {
			// Left half is sorted, search right half
			left = mid + 1
		} else {
			// Right half is sorted, search left half
			right = mid - 1
		}
	}
	
	return nums[left]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search for Pivot Point
- **Rotation Detection**: Compare first and last elements to check rotation
- **Pivot Finding**: Use binary search to find minimum (rotation point)
- **Neighbor Comparison**: Check if mid is pivot by comparing with neighbors
- **Range Reduction**: Eliminate half based on sorted property

## 2. PROBLEM CHARACTERISTICS
- **Rotated Sorted Array**: Originally sorted, then rotated at unknown pivot
- **No Duplicates**: All elements are unique
- **Minimum Search**: Find the smallest element (rotation point)
- **Logarithmic Time**: O(log N) solution required

## 3. SIMILAR PROBLEMS
- Search in Rotated Sorted Array (LeetCode 33)
- Find Minimum in Rotated Sorted Array II (with duplicates, LeetCode 154)
- Find Rotation Count
- Peak Finding (LeetCode 162)

## 4. KEY OBSERVATIONS
- **Pivot property**: Minimum is where array "wraps around"
- **Sorted halves**: At least one half is always sorted
- **Comparison logic**: Use nums[mid] with nums[left] to decide direction
- **Early exit**: If array not rotated, first element is minimum

## 5. VARIATIONS & EXTENSIONS
- **With Duplicates**: More complex, may need O(N) in worst case
- **Find Rotation Count**: Return index of minimum instead of value
- **Multiple Rotations**: Handle arrays rotated multiple times
- **Peak Finding**: Find maximum instead of minimum

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are there duplicates? How many rotations?"
- Edge cases: empty array, single element, no rotation
- Space complexity: O(1) - only pointers and variables
- Time complexity: O(log N) - binary search for pivot

## 7. COMMON MISTAKES
- Not handling case where array is not rotated
- Wrong comparison logic for deciding search direction
- Off-by-one errors in neighbor checking
- Using standard binary search without rotation handling
- Not handling single element arrays

## 8. OPTIMIZATION STRATEGIES
- **Early rotation check**: If nums[0] <= nums[n-1], array not rotated
- **Pivot detection**: Check if mid is pivot by comparing with neighbors
- **Template method**: Use consistent pivot search pattern
- **Boundary optimization**: Handle edge cases efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the break point in a circular list:**
- You have a sorted list that was cut and the pieces swapped
- The smallest element is where the list "wraps around"
- One side of any midpoint will always be sorted
- If the left side is sorted, the minimum must be in the unsorted right side
- Keep narrowing down until you find the exact break point

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Rotated sorted array with unique elements
2. **Goal**: Find the minimum element
3. **Output**: Value of minimum element
4. **Constraint**: Array was originally sorted, then rotated

#### Phase 2: Key Insight Recognition
- **"Pivot concept"** → Minimum is where rotation occurred
- **"Sorted half property"** → One half is always sorted
- **"Direction decision"** → Use nums[mid] with nums[left] to choose direction
- **"Early exit"** → If not rotated, first element is minimum

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find where the sorted array was rotated.
The minimum element is at the rotation point.
At any position, one half of the array will be sorted.
If the left half is sorted, the minimum must be in the right half.
If the left half contains the rotation, search there.
I'll keep narrowing until I find the exact minimum."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Not possible per problem constraints
- **Single element**: Return that element
- **No rotation**: First element is minimum
- **Full rotation**: Same as no rotation case

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums=[3,4,5,1,2]

Human thinking:
"I'll find the minimum in the rotated array:

Check if rotated: nums[0]=3 > nums[4]=2 → Yes, rotated

Initial: left=0, right=4
mid=2, nums[mid]=5
Check if mid is pivot: nums[2]=5 > nums[3]=1 → Yes! nums[3] is minimum
Return nums[3]=1

Another example: nums=[4,5,6,7,0,1,2,3]

Initial: left=0, right=7
mid=3, nums[mid]=7
nums[mid] > nums[left] (7 > 4) → left half sorted, search right half
left=4, right=7

mid=5, nums[mid]=1
nums[mid] > nums[left] (1 > 0) → left half unsorted, search left half
left=4, right=4

Return nums[4]=0"
```

#### Phase 6: Intuition Validation
- **Why pivot is minimum**: Rotation breaks the sorted order at minimum
- **Why one half sorted**: Rotation only affects one boundary
- **Why O(log N) time**: Still eliminate half elements each iteration
- **Why early exit works**: Unrotated array has minimum at start

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just find the smallest element?"** → That's O(N), need O(log N)
2. **"How do I know which half to search?"** → Compare nums[mid] with nums[left]
3. **"What if the array isn't rotated?"** → Check nums[0] <= nums[n-1] first
4. **"Can I use standard binary search?"** → No, need modified logic for rotation

### Real-World Analogy
**Like finding the starting point of a circular race track:**
- You have a circular track with distance markers in order
- The track was cut and the start/end point moved to a new position
- The smallest distance marker is where the track was cut
- At any point on the track, one direction goes to higher numbers
- The other direction goes to lower numbers (toward the start)
- Keep following the lower numbers until you find the smallest

### Human-Readable Pseudocode
```
function findMinInRotated(array):
    if length(array) == 1:
        return array[0]
    
    # Check if array is not rotated
    if array[0] <= array[length(array)-1]:
        return array[0]
    
    left = 0
    right = length(array) - 1
    
    while left < right:
        middle = left + (right - left) // 2
        
        # Check if middle is the pivot
        if middle > 0 and array[middle] < array[middle-1]:
            return array[middle]
        
        # Check if middle+1 is the pivot
        if middle < length(array)-1 and array[middle] > array[middle+1]:
            return array[middle+1]
        
        # Decide which half to search
        if array[middle] >= array[left]:
            # Left half is sorted, search right half
            left = middle + 1
        else:
            # Right half is sorted, search left half
            right = middle - 1
    
    return array[left]
```

### Execution Visualization

### Example: nums=[3,4,5,1,2]
```
Array: [3][4][5][1][2]
Index:   0  1  2  3  4

Check rotation: nums[0]=3 > nums[4]=2 → Rotated

Initial: left=0, right=4
mid=2, nums[mid]=5
Check if mid is pivot: nums[2]=5 > nums[3]=1 → Yes!
Minimum is nums[3]=1
```

### Example: nums=[4,5,6,7,0,1,2,3]
```
Array: [4][5][6][7][0][1][2][3]
Index:   0  1  2  3  4  5  6  7

Check rotation: nums[0]=4 > nums[7]=3 → Rotated

Initial: left=0, right=7
mid=3, nums[mid]=7
nums[mid] >= nums[left] (7 >= 4) → left half sorted, search right
left=4, right=7

mid=5, nums[mid]=1
nums[mid] < nums[left] (1 < 4) → search left half
right=4

mid=4, nums[mid]=0
nums[mid] < nums[left] (0 < 4) → search left half
right=3

left=4 > right=3 → Minimum is nums[4]=0
```

### Key Visualization Points:
- **Rotation check**: Compare first and last elements
- **Pivot detection**: Check if mid is smaller than previous element
- **Direction logic**: Use nums[mid] with nums[left] to decide
- **Convergence**: Search space narrows to single element

### Memory Layout Visualization:
```
Array: [4][5][6][7][0][1][2][3]
Index:   0  1  2  3  4  5  6  7
        ^           ^
     left=0      right=7
        mid=3 (value=7)
        left half [0-3] sorted: [4,5,6,7]
```

### Time Complexity Breakdown:
- **Each iteration**: O(1) - constant work
- **Number of iterations**: O(log N) - halving search space
- **Total time**: O(log N)
- **Space**: O(1) - only pointers and variables

### Alternative Approaches:

#### 1. Linear Scan (O(N) time, O(1) space)
```go
func findMin(nums []int) int {
    minVal := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < minVal {
            minVal = nums[i]
        }
    }
    return minVal
}
```
- **Pros**: Simple, works for any array
- **Cons**: O(N) time, doesn't leverage sorted property

#### 2. Find Rotation Count First (O(log N) time, O(1) space)
```go
func findMin(nums []int) int {
    // Find rotation index
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}
```
- **Pros**: Cleaner logic, same complexity
- **Cons**: Different approach, same result

#### 3. Recursive Version (O(log N) time, O(log N) space)
```go
func findMin(nums []int) int {
    return findMinHelper(nums, 0, len(nums)-1)
}

func findMinHelper(nums []int, left, right int) int {
    if left == right {
        return nums[left]
    }
    
    if nums[left] <= nums[right] {
        return nums[left]
    }
    
    mid := left + (right-left)/2
    if nums[mid] >= nums[left] {
        return findMinHelper(nums, mid+1, right)
    } else {
        return findMinHelper(nums, left, mid)
    }
}
```
- **Pros**: Elegant, same time complexity
- **Cons**: Uses O(log N) stack space

### Extensions for Interviews:
- **With Duplicates**: Handle arrays with duplicate elements
- **Find Rotation Count**: Return index of minimum instead of value
- **Find Maximum**: Similar logic for finding maximum in rotated array
- **Multiple Queries**: Handle multiple minimum queries efficiently
*/
func main() {
	// Test cases
	testCases := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{11, 13, 15, 17, 2, 5, 6, 8, 10},
		{1, 2, 3, 4, 5},
		{2, 1},
		{1},
		{5, 1, 2, 3, 4},
		{2, 3, 4, 5, 1},
		{3, 4, 5, 6, 7, 0, 1, 2},
		{10, 20, 30, 40, 5, 6, 7},
	}
	
	for i, nums := range testCases {
		result := findMin(nums)
		fmt.Printf("Test Case %d: %v -> Minimum: %d\n", i+1, nums, result)
	}
}
