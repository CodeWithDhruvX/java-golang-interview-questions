package main

import "fmt"

// 78. Subsets
// Time: O(N * 2^N), Space: O(N) for recursion + O(2^N) for result
func subsets(nums []int) [][]int {
	var result [][]int
	current := make([]int, 0, len(nums))
	
	backtrackSubsets(nums, 0, current, &result)
	return result
}

func backtrackSubsets(nums []int, start int, current []int, result *[][]int) {
	// Add current subset to result
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)
	
	// Generate subsets by including each remaining element
	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrackSubsets(nums, i+1, current, result)
		current = current[:len(current)-1]
	}
}

// Iterative approach
func subsetsIterative(nums []int) [][]int {
	result := [][]int{{}} // Start with empty subset
	
	for _, num := range nums {
		newSubsets := make([][]int, 0, len(result))
		
		// For each existing subset, create a new subset with current number
		for _, subset := range result {
			newSubset := make([]int, len(subset))
			copy(newSubset, subset)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		
		// Add all new subsets to result
		result = append(result, newSubsets...)
	}
	
	return result
}

// Bit manipulation approach
func subsetsBit(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<n)
	
	for mask := 0; mask < 1<<n; mask++ {
		subset := make([]int, 0)
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Backtracking for Power Set Generation
- **Include/Exclude Decision**: For each element, decide to include or exclude
- **Recursive Building**: Build subsets incrementally
- **Current Subset Tracking**: Maintain current subset during exploration
- **All Combinations**: Generate all 2^N possible subsets

## 2. PROBLEM CHARACTERISTICS
- **Power Set**: Generate all subsets of a set
- **Binary Choices**: Each element has 2 choices (include/exclude)
- **Order Independence**: Subsets are sets, order doesn't matter
- **Empty Set**: Always include empty subset as valid result

## 3. SIMILAR PROBLEMS
- Subsets II (LeetCode 90) - Handle duplicate elements
- Combination Sum (LeetCode 39) - Subsets with sum constraint
- Permutations (LeetCode 46) - All orderings instead of subsets
- Combination Sum III (LeetCode 216) - Fixed size subsets

## 4. KEY OBSERVATIONS
- **2^N subsets**: Each element doubles the number of subsets
- **Backtracking natural fit**: Need to explore all include/exclude combinations
- **Start parameter**: Ensures each subset is generated once
- **Empty subset**: Valid subset that should be included

## 5. VARIATIONS & EXTENSIONS
- **Duplicate Elements**: Handle repeated elements in input
- **Fixed Size**: Generate subsets of exactly k elements
- **Sum Constraint**: Generate subsets with specific sum
- **Lexicographic Order**: Generate subsets in sorted order

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are elements unique? What about empty input?"
- Edge cases: empty array, single element, duplicate elements
- Time complexity: O(N × 2^N) - 2^N subsets, each takes O(N) to copy
- Space complexity: O(N) for recursion + O(N × 2^N) for result

## 7. COMMON MISTAKES
- Not adding empty subset to result
- Making shallow copies of current subset
- Using wrong start parameter (causes duplicates)
- Not handling empty input case
- Forgetting to backtrack properly

## 8. OPTIMIZATION STRATEGIES
- **Iterative approach**: Avoid recursion stack
- **Bit manipulation**: Use binary representation for subset selection
- **Early termination**: Not applicable for complete power set
- **Memory optimization**: Reuse slices when possible

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding all possible teams from a group of people:**
- You have N people (array elements)
- You want to find all possible team combinations
- For each person, you decide whether to include them or not
- The empty team (no one) is also a valid team
- Continue until you've considered all people for all combinations

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of distinct elements
2. **Goal**: Generate all possible subsets (power set)
3. **Output**: List of all subsets, including empty set
4. **Constraint**: Order within subsets doesn't matter

#### Phase 2: Key Insight Recognition
- **"Binary decisions"** → Each element: include or exclude
- **"Power set size"** → 2^N subsets total
- **"Backtracking natural fit"** → Need to explore all combinations
- **"Start parameter"** → Prevent duplicate subsets

#### Phase 3: Strategy Development
```
Human thought process:
"I need to generate all subsets of the array.
For each element, I have two choices: include it or exclude it.
I'll use backtracking to explore all combinations.
At each step, I'll add the current subset to result.
Then I'll try including the next element, and backtrack to try excluding it."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return [[]] (only empty subset)
- **Single element**: Return [[], [element]]
- **Large arrays**: Handle recursion depth appropriately
- **Duplicate elements**: Need special handling (Subsets II)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1,2,3]

Human thinking:
"I'll build subsets by making include/exclude decisions:

Start with empty subset: []
Add [] to result

Consider element 1:
- Include 1: current=[1]
  Add [1] to result
  Consider element 2:
  - Include 2: current=[1,2]
    Add [1,2] to result
    Consider element 3:
    - Include 3: current=[1,2,3]
      Add [1,2,3] to result
    Backtrack: remove 3 → [1,2]
    - Exclude 3: current=[1,2]
      (already added)
  Backtrack: remove 2 → [1]
  - Exclude 2: current=[1]
    Consider element 3:
    - Include 3: current=[1,3]
      Add [1,3] to result
    Backtrack: remove 3 → [1]
    - Exclude 3: current=[1]
      (already added)
  Backtrack: remove 1 → []

- Exclude 1: current=[]
  Consider element 2:
  - Include 2: current=[2]
    Add [2] to result
    Consider element 3:
    - Include 3: current=[2,3]
      Add [2,3] to result
    Backtrack: remove 3 → [2]
    - Exclude 3: current=[2]
      (already added)
  Backtrack: remove 2 → []
  - Exclude 2: current=[]
    Consider element 3:
    - Include 3: current=[3]
      Add [3] to result
    Backtrack: remove 3 → []
    - Exclude 3: current=[]
      (already added)

All subsets: [], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]"
```

#### Phase 6: Intuition Validation
- **Why backtracking works**: Need to explore all include/exclude combinations
- **Why 2^N subsets**: Each element doubles the number of possibilities
- **Why start parameter**: Prevents generating same subset multiple times
- **Why O(N × 2^N) time**: 2^N subsets, each takes O(N) to copy

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → Need variable number of loops, backtracking is better
2. **"Should I sort first?"** → Not necessary for basic subsets, but helps for variations
3. **"What about very large N?"** → Complexity is inherently exponential
4. **"Can I optimize space?"** → Use bit manipulation or iterative approach

### Real-World Analogy
**Like finding all possible meal combinations from a menu:**
- You have N menu items (array elements)
- You want to find all possible meal combinations
- For each item, you decide whether to order it or not
- The empty meal (ordering nothing) is also a valid choice
- Continue until you've considered all items for all combinations

### Human-Readable Pseudocode
```
function subsets(nums):
    result = []
    current = []
    
    backtrack(nums, 0, current, result)
    return result

function backtrack(nums, start, current, result):
    // Add current subset to result
    add copy of current to result
    
    // Generate subsets by including each remaining element
    for i from start to length(nums)-1:
        current.append(nums[i])
        backtrack(nums, i + 1, current, result)  // i+1 to avoid duplicates
        current.pop()  // Backtrack
```

### Execution Visualization

### Example: nums = [1,2,3]
```
Recursion Tree:
                    [], start=0
                   /         \
                [1]           [] (exclude 1)
                start=1       start=1
               /     \         /     \
            [1,2]   [1]      [2]     []
            start=2  start=2  start=2  start=2
           /   \      / \      / \      / \
       [1,2,3] [1,2] [1,3] [1] [2,3] [2] [3] []
       (add)   (add) (add) (add) (add) (add) (add) (add)

All subsets added at each node: [], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]
```

### Key Visualization Points:
- **Include/exclude decisions**: Each level represents decision for one element
- **Current subset tracking**: Build subset incrementally
- **Result addition**: Add current subset at each node
- **Backtracking**: Remove last element when exploring different path
- **Complete coverage**: All 2^N subsets generated exactly once

### Memory Layout Visualization:
```
Current Subset Evolution:
[] → [1] → [1,2] → [1,2,3] ✓
          → [1,3] ✓
     → [2] → [2,3] ✓
          → [3] ✓

Result Building:
[] (initial)
→ [] (at root)
→ [1] (after including 1)
→ [1,2] (after including 2)
→ [1,2,3] (after including 3)
→ [1,3] (backtrack, include 3)
→ [2] (backtrack to root, include 2)
→ [2,3] (include 3)
→ [3] (backtrack to root, include 3)
```

### Time Complexity Breakdown:
- **Number of subsets**: 2^N (power set)
- **Copy time per subset**: O(N) to copy current subset
- **Total time**: O(N × 2^N)
- **Space complexity**: O(N) for recursion + O(N × 2^N) for result
- **Current subset**: O(N) additional space

### Alternative Approaches:

#### 1. Iterative Approach (O(N × 2^N) time, O(2^N) space)
```go
func subsetsIterative(nums []int) [][]int {
    result := [][]int{{}} // Start with empty subset
    
    for _, num := range nums {
        newSubsets := make([][]int, 0, len(result))
        
        // For each existing subset, create a new subset with current number
        for _, subset := range result {
            newSubset := make([]int, len(subset))
            copy(newSubset, subset)
            newSubset = append(newSubset, num)
            newSubsets = append(newSubsets, newSubset)
        }
        
        // Add all new subsets to result
        result = append(result, newSubsets...)
    }
    
    return result
}
```
- **Pros**: Iterative, no recursion stack
- **Cons**: More memory for intermediate results

#### 2. Bit Manipulation (O(N × 2^N) time, O(2^N) space)
```go
func subsetsBit(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0, 1<<n)
    
    for mask := 0; mask < 1<<n; mask++ {
        subset := make([]int, 0)
        for i := 0; i < n; i++ {
            if (mask>>i)&1 == 1 {
                subset = append(subset, nums[i])
            }
        }
        result = append(result, subset)
    }
    
    return result
}
```
- **Pros**: Elegant, uses binary representation
- **Cons**: Less intuitive, requires bit operations

#### 3. DFS with Fixed Size (O(N × C(N,k)) time, O(N) space)
```go
func subsetsFixedSize(nums []int, k int) [][]int {
    var result [][]int
    current := make([]int, 0, k)
    
    backtrackFixedSize(nums, 0, k, current, &result)
    return result
}

func backtrackFixedSize(nums []int, start, k int, current []int, result *[][]int) {
    if len(current) == k {
        temp := make([]int, k)
        copy(temp, current)
        *result = append(*result, temp)
        return
    }
    
    for i := start; i < len(nums); i++ {
        current = append(current, nums[i])
        backtrackFixedSize(nums, i+1, k, current, result)
        current = current[:len(current)-1]
    }
}
```
- **Pros**: Generates subsets of fixed size only
- **Cons**: Limited to specific subset sizes

### Extensions for Interviews:
- **Subsets II**: Handle duplicate elements in input
- **Combination Sum**: Generate subsets with sum constraint
- **Combination Sum III**: Generate subsets of exactly k numbers summing to target
- **Lexicographic Order**: Generate subsets in sorted order
- **Large Input**: Handle very large arrays efficiently
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3},
		{0},
		{},
		{1, 2},
		{1, 2, 3, 4},
		{5},
		{1, 1, 2},
		{-1, -2, -3},
		{100, 200, 300},
		{1},
	}
	
	for i, nums := range testCases {
		result1 := subsets(nums)
		result2 := subsetsIterative(nums)
		result3 := subsetsBit(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Recursive: %d subsets\n", len(result1))
		fmt.Printf("  Iterative: %d subsets\n", len(result2))
		fmt.Printf("  Bit method: %d subsets\n\n", len(result3))
	}
}
