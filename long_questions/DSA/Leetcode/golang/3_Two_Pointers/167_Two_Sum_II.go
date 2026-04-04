package main

import "fmt"

// 167. Two Sum II - Input Array Is Sorted
// Time: O(N), Space: O(1)
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	
	for left < right {
		sum := numbers[left] + numbers[right]
		
		if sum == target {
			return []int{left + 1, right + 1} // 1-indexed
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	
	return []int{}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers on Sorted Array
- **Left Pointer**: Starts at beginning, moves rightward
- **Right Pointer**: Starts at end, moves leftward
- **Sum Comparison**: Compare sum with target
- **Directional Movement**: Move pointers based on sum comparison

## 2. PROBLEM CHARACTERISTICS
- **Sorted Input**: Array is already sorted in ascending order
- **Exactly One Solution**: Guaranteed to have exactly one valid pair
- **1-Indexed Output**: Return indices starting from 1, not 0
- **No Reuse**: Cannot use same element twice

## 3. SIMILAR PROBLEMS
- Two Sum (LeetCode 1) - unsorted array, uses hash map
- 3Sum (LeetCode 15) - extends to three numbers
- Two Sum II - Input array is sorted (current problem)
- Two Sum III - Data structure design

## 4. KEY OBSERVATIONS
- **Sorted advantage**: Can use two-pointer approach instead of hash map
- **Monotonic property**: Moving left increases sum, moving right decreases sum
- **Space efficiency**: O(1) space vs O(N) for hash map approach
- **Guaranteed solution**: Don't need to handle no-solution case

## 5. VARIATIONS & EXTENSIONS
- **Multiple solutions**: Return all valid pairs
- **K-Sum**: Extend to K numbers
- **Approximate target**: Find closest sum to target
- **Descending order**: Handle arrays sorted in reverse

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is array sorted? Are indices 1-based or 0-based?"
- Edge cases: empty array, single element, negative numbers
- Space complexity: O(1) - only two pointers
- Time complexity: O(N) - single pass with two pointers

## 7. COMMON MISTAKES
- Using hash map approach when two-pointer is better
- Forgetting 1-based indexing requirement
- Moving both pointers in wrong direction
- Not handling negative numbers correctly
- Not using sorted property to advantage

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(1) space
- **Early termination**: Can stop when pointers cross
- **Cache optimization**: Sequential memory access pattern
- **No additional data structures**: Pure pointer manipulation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding two numbers that add up to a target in a phone book:**
- The phone book is sorted alphabetically (numbers are sorted)
- You want to find two people whose phone numbers sum to a target
- Start with the first person and the last person
- If their sum is too small, move to the next person
- If their sum is too large, move to the previous person
- Continue until you find the perfect pair

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted array of numbers and a target sum
2. **Goal**: Find two numbers that sum to target
3. **Output**: 1-based indices of the two numbers
4. **Constraint**: Exactly one solution exists

#### Phase 2: Key Insight Recognition
- **"Sorted advantage"** → Can use two-pointer instead of hash map
- **"Monotonic movement"** → Moving pointers changes sum predictably
- **"Space efficiency"** → No need for extra storage
- **"1-based indexing"** → Important output format detail

#### Phase 3: Strategy Development
```
Human thought process:
"The array is sorted, so I can use two pointers efficiently.
I'll start with the first and last elements.
If their sum is too small, I need a larger sum, so I'll move the left pointer right.
If their sum is too large, I need a smaller sum, so I'll move the right pointer left.
Since the array is sorted, this guarantees I'll find the unique solution.
I'll return 1-based indices as required."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty (though problem guarantees solution)
- **Single element**: Return empty (need two elements)
- **Negative numbers**: Handled correctly by sum comparison
- **Large arrays**: Two-pointer approach remains efficient

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [2, 7, 11, 15], target = 9
left = 0 (value=2), right = 3 (value=15)

Human thinking:
"Let's try the outermost pair:
Sum = 2 + 15 = 17
This is greater than target (9), so I need a smaller sum.
I'll move the right pointer left to get a smaller number.

left = 0 (2), right = 2 (11):
Sum = 2 + 11 = 13
Still greater than 9, move right left again.

left = 0 (2), right = 1 (7):
Sum = 2 + 7 = 9
Perfect! This equals the target.
Return 1-based indices: [1, 2]"
```

#### Phase 6: Intuition Validation
- **Why two pointers work**: Sorted array gives monotonic sum changes
- **Why it's efficient**: Each step eliminates one possibility
- **Why it's correct**: Guaranteed unique solution ensures we'll find it
- **Why O(1) space**: No need for hash map or extra storage

### Common Human Pitfalls & How to Avoid Them
1. **"Should I use hash map?"** → No, two-pointer is better for sorted arrays
2. **"What about 0-based vs 1-based?"** → Problem requires 1-based output
3. **"Why not try all pairs?"** → That's O(N²), two-pointer is O(N)
4. **"Can pointers move wrong way?"** → Left only right, right only left

### Real-World Analogy
**Like finding two ingredients that weigh exactly target amount:**
- You have ingredients sorted by weight
- You need two that sum to exact target weight
- Start with lightest and heaviest
- If too heavy, replace heaviest with next lighter
- If too light, replace lightest with next heavier
- Continue until perfect combination found

### Human-Readable Pseudocode
```
function twoSumSorted(numbers, target):
    left = 0
    right = length(numbers) - 1
    
    while left < right:
        currentSum = numbers[left] + numbers[right]
        
        if currentSum == target:
            return [left + 1, right + 1]  // 1-based
        else if currentSum < target:
            left = left + 1  // Need larger sum
        else:
            right = right - 1  // Need smaller sum
    
    return []  // No solution (shouldn't happen per problem)
```

### Execution Visualization

### Example: [2, 7, 11, 15], target = 9
```
Array: [2, 7, 11, 15]
Index:  0, 1,  2,  3

Initial: left=0(2), right=3(15)
Sum = 2 + 15 = 17 > 9
Move right left: right=2

Step 2: left=0(2), right=2(11)
Sum = 2 + 11 = 13 > 9
Move right left: right=1

Step 3: left=0(2), right=1(7)
Sum = 2 + 7 = 9 == target
Found! Return [1, 2] (1-based)
```

### Example with negative numbers: [-10, -5, -3, 0, 1, 3, 5, 10], target = 0
```
Array: [-10, -5, -3, 0, 1, 3, 5, 10]
Index:   0,  1,  2, 3, 4, 5, 6, 7

Initial: left=0(-10), right=7(10)
Sum = -10 + 10 = 0 == target
Found! Return [1, 8] (1-based)
```

### Key Visualization Points:
- **Left pointer**: Only moves right (increases sum)
- **Right pointer**: Only moves left (decreases sum)
- **Sum comparison**: Determines which pointer to move
- **1-based output**: Add 1 to indices before returning

### Memory Layout Visualization:
```
Array: [2][7][11][15]
Index:  0  1  2  3
        ^        ^
     left=0   right=3
     sum=17 > 9

        ^     ^
     left=0   right=1
     sum=9 == target
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited at most once
- **Constant Space**: O(1) - only two pointers
- **No Additional Data Structures**: Pure pointer manipulation
- **Early Termination**: Returns as soon as solution found

### Alternative Approaches:

#### 1. Hash Map Approach (O(N) time, O(N) space)
```go
func twoSumII(numbers []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range numbers {
        complement := target - num
        if j, exists := numMap[complement]; exists {
            return []int{j + 1, i + 1} // 1-based
        }
        numMap[num] = i
    }
    return []int{}
}
```
- **Pros**: Works for unsorted arrays too
- **Cons**: Uses O(N) extra space, ignores sorted property

#### 2. Binary Search Approach (O(N log N) time, O(1) space)
```go
func twoSumII(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        complement := target - numbers[i]
        if j := binarySearch(numbers, complement, i+1); j != -1 {
            return []int{i + 1, j + 1}
        }
    }
    return []int{}
}

func binarySearch(nums []int, target, start int) int {
    left, right := start, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
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
- **Pros**: O(1) space
- **Cons**: O(N log N) time, slower than two-pointer

### Comparison with Two Sum (LeetCode 1):
| Aspect | Two Sum (unsorted) | Two Sum II (sorted) |
|--------|-------------------|---------------------|
| Approach | Hash Map | Two Pointers |
| Time | O(N) | O(N) |
| Space | O(N) | O(1) |
| Input | Unsorted | Sorted |
| Indexing | 0-based | 1-based |
| Solution | Not guaranteed | Exactly one |

### Extensions for Interviews:
- **Multiple pairs**: Find all pairs that sum to target
- **Closest sum**: Find pair with sum closest to target
- **K-Sum**: Extend to K numbers using recursion
- **Descending order**: Handle arrays sorted in reverse
*/
func main() {
	// Test cases
	testCases := []struct {
		numbers []int
		target  int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 3, 4}, 6},
		{[]int{-1, 0}, -1},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8},
		{[]int{-10, -5, -3, 0, 1, 3, 5, 10}, 0},
		{[]int{1, 3, 5, 7, 9}, 12},
	}
	
	for i, tc := range testCases {
		result := twoSumII(tc.numbers, tc.target)
		fmt.Printf("Test Case %d: numbers=%v, target=%d -> Indices: %v\n", 
			i+1, tc.numbers, tc.target, result)
	}
}
