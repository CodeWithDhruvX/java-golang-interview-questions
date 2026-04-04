package main

import "fmt"

// 46. Permutations
// Time: O(N * N!), Space: O(N!) for result + O(N) for recursion
func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	current := make([]int, len(nums))
	
	backtrackPermute(nums, used, current, 0, &result)
	return result
}

func backtrackPermute(nums []int, used []bool, current []int, pos int, result *[][]int) {
	if pos == len(nums) {
		// Make a copy of current permutation
		temp := make([]int, len(nums))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			current[pos] = nums[i]
			
			backtrackPermute(nums, used, current, pos+1, result)
			
			used[i] = false
		}
	}
}

// Alternative approach using swapping
func permuteSwap(nums []int) [][]int {
	var result [][]int
	backtrackSwap(nums, 0, &result)
	return result
}

func backtrackSwap(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// Make a copy of current permutation
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrackSwap(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Backtracking for Permutations
- **Position-based Selection**: Fill each position with unused elements
- **Used Tracking**: Track which elements have been used in current permutation
- **Recursive Building**: Build permutations position by position
- **Backtracking**: Remove choices when exploring different paths

## 2. PROBLEM CHARACTERISTICS
- **Permutation Generation**: Generate all possible orderings of elements
- **Unique Elements**: Input contains distinct elements
- **Complete Permutations**: Each permutation uses all elements exactly once
- **Order Matters**: Different orderings are distinct permutations

## 3. SIMILAR PROBLEMS
- Permutations II (LeetCode 47) - Handle duplicate elements
- Next Permutation (LeetCode 31) - Find next lexicographic permutation
- Permutation Sequence (LeetCode 60) - Find kth permutation
- Subsets (LeetCode 78) - Generate all subsets

## 4. KEY OBSERVATIONS
- **Backtracking natural fit**: Need to explore all possible arrangements
- **Used array essential**: Track which elements are available for current position
- **Position parameter**: Track current position being filled
- **Complete permutation**: Base case when all positions are filled

## 5. VARIATIONS & EXTENSIONS
- **Duplicate Elements**: Handle repeated elements in input
- **Partial Permutations**: Generate permutations of length k
- **Lexicographic Order**: Generate permutations in sorted order
- **Kth Permutation**: Find specific permutation without generating all

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are elements unique? What about empty input?"
- Edge cases: empty array, single element, large arrays
- Time complexity: O(N × N!) - N! permutations, each takes O(N) to copy
- Space complexity: O(N) for recursion + O(N × N!) for result

## 7. COMMON MISTAKES
- Not using used array (causes duplicates in permutations)
- Making shallow copies of current permutation
- Not resetting used array properly during backtracking
- Using wrong base case condition
- Not handling empty input case

## 8. OPTIMIZATION STRATEGIES
- **Swap method**: More memory efficient, no extra used array
- **In-place modification**: Modify array directly instead of using extra space
- **Early pruning**: Not applicable for complete permutations
- **Iterative generation**: Use Heap's algorithm for iterative approach

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like arranging books on a shelf:**
- You have N distinct books (array elements)
- You want to find all possible ways to arrange them
- For each position, you choose from the remaining books
- Once you place a book, you can't use it again in current arrangement
- Continue until all positions are filled, then backtrack to try different arrangements

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of distinct elements
2. **Goal**: Generate all possible permutations (orderings)
3. **Output**: List of all permutations, each containing all elements
4. **Constraint**: Each element used exactly once per permutation

#### Phase 2: Key Insight Recognition
- **"Backtracking natural fit"** → Need to explore all possible arrangements
- **"Position-based selection"** → Fill each position with available elements
- **"Used tracking essential"** → Need to track which elements are used
- **"Complete coverage"** → Each permutation must use all elements

#### Phase 3: Strategy Development
```
Human thought process:
"I need to generate all permutations of the array.
I'll use backtracking to build permutations position by position.
For each position, I'll try each unused element.
I'll mark elements as used when I include them.
When I reach the end, I have a complete permutation.
Then I backtrack and try different choices."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty result
- **Single element**: Return [[element]]
- **Large arrays**: Handle recursion depth and memory appropriately
- **Duplicate elements**: Need special handling (Permutations II)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1,2,3]

Human thinking:
"I'll build permutations position by position:

Position 0:
- Try 1 (unused): current=[1,_ ,_]
  Mark 1 as used
  Position 1:
  - Try 2 (unused): current=[1,2,_]
    Mark 2 as used
    Position 2:
    - Try 3 (unused): current=[1,2,3] → COMPLETE [1,2,3]
    Backtrack: unmark 3
  Backtrack: unmark 2
  - Try 3 (unused): current=[1,3,_]
    Mark 3 as used
    Position 2:
    - Try 2 (unused): current=[1,3,2] → COMPLETE [1,3,2]
    Backtrack: unmark 2
  Backtrack: unmark 3
Backtrack: unmark 1

Position 0:
- Try 2 (unused): current=[2,_ ,_]
  Mark 2 as used
  Position 1:
  - Try 1 (unused): current=[2,1,_]
    Mark 1 as used
    Position 2:
    - Try 3 (unused): current=[2,1,3] → COMPLETE [2,1,3]
  Backtrack: unmark 1
  - Try 3 (unused): current=[2,3,_]
    Mark 3 as used
    Position 2:
    - Try 1 (unused): current=[2,3,1] → COMPLETE [2,3,1]
  Backtrack: unmark 3
Backtrack: unmark 2

Position 0:
- Try 3 (unused): current=[3,_ ,_]
  Similar process generates [3,1,2] and [3,2,1]

All permutations: [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]"
```

#### Phase 6: Intuition Validation
- **Why backtracking works**: Need to explore all possible arrangements
- **Why used array works**: Prevents reusing elements in same permutation
- **Why O(N × N!) time**: N! permutations, each takes O(N) to copy
- **Why position-based works**: Natural way to build permutations

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use next_permutation?"** → That finds the next one, not all permutations
2. **"Should I use recursion or iteration?"** → Recursion is more intuitive for this problem
3. **"What about very large N?"** → Complexity is inherently factorial, can't do much better
4. **"Can I optimize space?"** → Use swap method to avoid used array

### Real-World Analogy
**Like finding all possible seating arrangements:**
- You have N distinct people (array elements)
- You need to find all possible ways to seat them in N chairs
- For each chair, you choose from the remaining unseated people
- Once someone is seated, they can't be seated again in current arrangement
- Continue until all chairs are filled, then try different arrangements

### Human-Readable Pseudocode
```
function permute(nums):
    result = []
    used = [false] * length(nums)
    current = [0] * length(nums)
    
    backtrack(nums, used, current, 0, result)
    return result

function backtrack(nums, used, current, pos, result):
    if pos == length(nums):
        add copy of current to result
        return
    
    for i from 0 to length(nums)-1:
        if not used[i]:
            used[i] = true
            current[pos] = nums[i]
            
            backtrack(nums, used, current, pos + 1, result)
            
            used[i] = false  // Backtrack
```

### Execution Visualization

### Example: nums = [1,2,3]
```
Recursion Tree (Used Array Method):
                    [], pos=0
                   /    |    \
                [1]    [2]    [3]
                pos=1   pos=1   pos=1
               / |     / |     / |
            [1,2] [1,3] [2,1] [2,3] [3,1] [3,2]
            pos=2   pos=2   pos=2   pos=2   pos=2   pos=2
             |      |      |      |      |      |
           [1,2,3] [1,3,2] [2,1,3] [2,3,1] [3,1,2] [3,2,1]

All permutations found at pos=3 (base case)
```

### Key Visualization Points:
- **Position-based building**: Each level fills one position
- **Used tracking**: Elements marked as used when included
- **Backtracking**: Elements unmarked when exploring different paths
- **Complete permutations**: Base case when all positions filled
- **Systematic exploration**: Try all unused elements at each position

### Memory Layout Visualization:
```
Used Array Evolution:
[false,false,false] → [true,false,false] → [true,true,false] → [true,true,true]
                      ↘ [true,false,true] → [true,false,true]
                      ↘ [false,true,false] → [false,true,true]
                      ↘ [false,false,true] → [false,true,true]

Current Array Evolution:
[_,_,_] → [1,_,_] → [1,2,_] → [1,2,3] ✓
         → [1,_,_] → [1,3,_] → [1,3,2] ✓
         → [2,_,_] → [2,1,_] → [2,1,3] ✓
         → [2,_,_] → [2,3,_] → [2,3,1] ✓
         → [3,_,_] → [3,1,_] → [3,1,2] ✓
         → [3,_,_] → [3,2,_] → [3,2,1] ✓
```

### Time Complexity Breakdown:
- **Number of permutations**: N! (factorial)
- **Copy time per permutation**: O(N) to copy current permutation
- **Total time**: O(N × N!)
- **Space complexity**: O(N) for recursion stack + O(N × N!) for result
- **Used array**: O(N) additional space

### Alternative Approaches:

#### 1. Swap Method (O(N × N!) time, O(N) space)
```go
func permuteSwap(nums []int) [][]int {
    var result [][]int
    backtrackSwap(nums, 0, &result)
    return result
}

func backtrackSwap(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        // Make a copy of current permutation
        temp := make([]int, len(nums))
        copy(temp, nums)
        *result = append(*result, temp)
        return
    }
    
    for i := start; i < len(nums); i++ {
        nums[start], nums[i] = nums[i], nums[start]
        backtrackSwap(nums, start+1, result)
        nums[start], nums[i] = nums[i], nums[start]
    }
}
```
- **Pros**: No extra used array, more memory efficient
- **Cons**: Modifies input array, need careful swapping

#### 2. Iterative Heap's Algorithm (O(N × N!) time, O(N) space)
```go
func permuteHeap(nums []int) [][]int {
    var result [][]int
    n := len(nums)
    c := make([]int, n) // Control array
    
    // Add initial permutation
    temp := make([]int, n)
    copy(temp, nums)
    result = append(result, temp)
    
    i := 0
    for i < n {
        if c[i] < i {
            if i%2 == 0 {
                nums[0], nums[i] = nums[i], nums[0]
            } else {
                nums[c[i]], nums[i] = nums[i], nums[c[i]]
            }
            
            temp := make([]int, n)
            copy(temp, nums)
            result = append(result, temp)
            
            c[i]++
            i = 0
        } else {
            c[i] = 0
            i++
        }
    }
    
    return result
}
```
- **Pros**: Iterative, no recursion stack
- **Cons**: Complex to understand and implement

#### 3. Lexicographic Generation (O(N × N!) time, O(1) extra space)
```go
func permuteLexicographic(nums []int) [][]int {
    sort.Ints(nums)
    var result [][]int
    
    // Add initial permutation
    temp := make([]int, len(nums))
    copy(temp, nums)
    result = append(result, temp)
    
    for nextPermutation(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        result = append(result, temp)
    }
    
    return result
}

func nextPermutation(nums []int) bool {
    n := len(nums)
    
    // Find the first element from the right that is smaller than the next element
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    
    if i < 0 {
        return false // Already the last permutation
    }
    
    // Find the smallest element greater than nums[i] from the right
    j := n - 1
    for nums[j] <= nums[i] {
        j--
    }
    
    // Swap them
    nums[i], nums[j] = nums[j], nums[i]
    
    // Reverse the suffix
    reverse(nums, i+1, n-1)
    
    return true
}

func reverse(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```
- **Pros**: Generates permutations in lexicographic order
- **Cons**: More complex implementation, requires sorting first

### Extensions for Interviews:
- **Permutations II**: Handle duplicate elements in input
- **Next Permutation**: Find next lexicographic permutation
- **Kth Permutation**: Find kth permutation without generating all
- **Partial Permutations**: Generate permutations of length k
- **Circular Permutations**: Generate circular arrangements
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3},
		{0, 1},
		{1},
		{},
		{1, 2, 3, 4},
		{1, 1, 2},
		{1, 2, 2},
		{5, 6, 7},
		{1, 2, 3, 4, 5},
	}
	
	for i, nums := range testCases {
		// Make copies for both methods
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		
		result1 := permute(nums1)
		result2 := permuteSwap(nums2)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Used array: %d permutations\n", len(result1))
		fmt.Printf("  Swap method: %d permutations\n\n", len(result2))
	}
}
