package main

import "fmt"

// 496. Next Greater Element I
// Time: O(N), Space: O(N)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// First pass: find next greater for all elements in nums2
	nextGreater := make(map[int]int)
	stack := []int{}
	
	for _, num := range nums2 {
		// While stack is not empty and current element is greater than stack top
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}
	
	// Elements remaining in stack have no next greater element
	for _, num := range stack {
		nextGreater[num] = -1
	}
	
	// Second pass: build result for nums1
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Monotonic Stack for Next Greater Element
- **Monotonic Decreasing Stack**: Maintain stack of decreasing elements
- **Next Greater Mapping**: Map each element to its next greater element
- **Two-Pass Solution**: First pass builds mapping, second pass queries results
- **Hash Map Lookup**: O(1) lookup for nums1 elements

## 2. PROBLEM CHARACTERISTICS
- **Next Greater Element**: Find first greater element to the right
- **Subset Query**: nums1 is subset of nums2
- **Order Preservation**: Original order in nums2 must be considered
- **Mapping Requirement**: Need to map nums2 elements to their next greater

## 3. SIMILAR PROBLEMS
- Next Greater Element II (LeetCode 503) - Circular array version
- Next Greater Element III (LeetCode 556) - Number digits version
- Daily Temperatures (LeetCode 739) - Similar next greater pattern
- Largest Rectangle in Histogram (LeetCode 84) - Monotonic stack application

## 4. KEY OBSERVATIONS
- **Stack property**: Stack maintains decreasing elements
- **Mapping strategy**: Build complete mapping for nums2 first
- **Efficient lookup**: Use hash map for O(1) queries
- **Two-pass approach**: Separate computation and query phases

## 5. VARIATIONS & EXTENSIONS
- **Circular array**: Wrap around to beginning of array
- **Next smaller element**: Find first smaller element to the right
- **K-th greater element**: Find k-th greater element
- **Multiple queries**: Answer queries for multiple subsets

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are elements unique? What about duplicate values?"
- Edge cases: empty arrays, no greater elements, all decreasing
- Time complexity: O(N + M) time, O(N) space where N=len(nums2), M=len(nums1)
- Space optimization: Can process nums1 during first pass

## 7. COMMON MISTAKES
- Not handling elements with no next greater (should be -1)
- Using nested loops (O(N×M) time)
- Not handling duplicate elements correctly
- Forgetting to process remaining stack elements
- Using two separate passes when one would suffice

## 8. OPTIMIZATION STRATEGIES
- **Monotonic stack**: O(N + M) time, O(N) space
- **Single pass**: Can process nums1 during first pass
- **Hash map lookup**: O(1) query time
- **Early termination**: Not applicable (need full processing)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the next taller person in a line:**
- You have people of different heights standing in a line (nums2)
- For each person, you want to know who's the first taller person to their right
- You scan from left to right, keeping track of people who haven't found their taller person yet
- When you find someone taller, you can resolve the shorter people behind them
- You record these relationships to answer questions about specific people (nums1)

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two arrays nums1 (subset) and nums2 (superset)
2. **Goal**: For each element in nums1, find next greater element in nums2
3. **Constraint**: Must be first greater element to the right
4. **Output**: Array of next greater elements or -1

#### Phase 2: Key Insight Recognition
- **"Monotonic stack natural fit"** → Maintain decreasing elements
- **"Next greater mapping"** → Build complete mapping first
- **"Hash map lookup"** → O(1) queries for nums1 elements
- **"Two-pass efficiency"** → Separate computation and query phases

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find next greater elements for nums1 elements in nums2.
I'll first process nums2 completely to build a mapping of each element to its next greater.
I'll use a monotonic decreasing stack to track elements waiting for their next greater.
When I find a greater element, I resolve all smaller elements in the stack.
Then I'll simply look up the results for nums1 elements using the mapping."
```

#### Phase 4: Edge Case Handling
- **Empty nums1**: Return empty array
- **Empty nums2**: Return array of -1s
- **No greater element**: Map to -1
- **Duplicate elements**: Handle correctly with stack

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums1 = [4, 1, 2], nums2 = [1, 3, 4, 2]

Human thinking:
"First, I'll process nums2 to build next greater mapping:

Initialize stack = [], nextGreater = {}

Process 1: stack = [1]
Process 3: 3 > stack.top=1, pop 1, nextGreater[1] = 3, stack = []
           push 3, stack = [3]
Process 4: 4 > stack.top=3, pop 3, nextGreater[3] = 4, stack = []
           push 4, stack = [4]
Process 2: 2 < stack.top=4, push 2, stack = [4, 2]

End of array: process remaining stack
pop 2: nextGreater[2] = -1
pop 4: nextGreater[4] = -1

Final mapping: {1:3, 3:4, 4:-1, 2:-1}

Now query nums1 elements:
4 → nextGreater[4] = -1
1 → nextGreater[1] = 3
2 → nextGreater[2] = -1

Result: [-1, 3, -1] ✓"
```

#### Phase 6: Intuition Validation
- **Why monotonic stack works**: Maintains elements waiting for greater element
- **Why mapping works**: Allows O(1) lookup for nums1 queries
- **Why two-pass efficient**: Separates computation from queries
- **Why O(N + M) time**: Each element processed once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N×M) time, too slow for large inputs
2. **"Should I process nums1 during first pass?"** → Yes, can optimize to single pass
3. **"What about duplicates?"** → Stack handles duplicates correctly
4. **"Can I optimize further?"** → Single pass is optimal

### Real-World Analogy
**Like finding the next higher salary in a company:**
- You have employees with different salaries listed in order of hiring
- For specific employees (nums1), you want to know who's the next person hired with higher salary
- You scan through the hiring list, keeping track of employees waiting to find someone with higher salary
- When you find someone with higher salary, you can resolve the lower-paid employees behind them
- You record these relationships to answer questions about specific employees

### Human-Readable Pseudocode
```
function nextGreaterElement(nums1, nums2):
    nextGreater = map()
    stack = []
    
    // Build mapping for nums2
    for num in nums2:
        while stack not empty and num > stack.top:
            top = stack.pop()
            nextGreater[top] = num
        stack.push(num)
    
    // Remaining elements have no next greater
    while stack not empty:
        top = stack.pop()
        nextGreater[top] = -1
    
    // Build result for nums1
    result = []
    for num in nums1:
        result.append(nextGreater[num])
    
    return result
```

### Execution Visualization

### Example: nums1 = [4, 1, 2], nums2 = [1, 3, 4, 2]
```
Stack Evolution and Mapping:
Initialize: stack = [], nextGreater = {}

Process 1: stack = [1]
Process 3: 3 > 1, pop 1, nextGreater[1] = 3, stack = []
           push 3, stack = [3]
Process 4: 4 > 3, pop 3, nextGreater[3] = 4, stack = []
           push 4, stack = [4]
Process 2: 2 < 4, push 2, stack = [4, 2]

Final Processing:
pop 2: nextGreater[2] = -1
pop 4: nextGreater[4] = -1

Final mapping: {1:3, 3:4, 4:-1, 2:-1}

Query nums1:
4 → -1
1 → 3
2 → -1

Result: [-1, 3, -1]
```

### Key Visualization Points:
- **Stack invariant**: Always maintains decreasing elements
- **Mapping building**: Pop when greater element found
- **Query phase**: Simple hash map lookup
- **No greater handling**: Map to -1 for remaining elements

### Memory Layout Visualization:
```
nums2: [1, 3, 4, 2]
       ↑  ↑  ↑  ↑
       1  3  4  2

Stack at position 3: [4, 2]
                    ↑  ↑
                    4  2 (values)

When processing 2:
2 < 4, so 2 is pushed
Stack becomes [4, 2]

Mapping after processing:
1 → 3 (first greater to right)
3 → 4 (first greater to right)
4 → -1 (no greater to right)
2 → -1 (no greater to right)
```

### Time Complexity Breakdown:
- **Monotonic stack**: O(N + M) time, O(N) space
- **Nested loops**: O(N × M) time, O(1) space (too slow)
- **Binary search**: O(N log N + M log N) time, O(N) space (more complex)
- **Brute force**: O(N × M) time, O(1) space

### Alternative Approaches:

#### 1. Single Pass Optimization (O(N + M) time, O(N) space)
```go
func nextGreaterElementOptimized(nums1 []int, nums2 []int) []int {
    nextGreater := make(map[int]int)
    stack := []int{}
    result := make([]int, len(nums1))
    
    // Build index mapping for nums1
    indexMap := make(map[int]int)
    for i, num := range nums1 {
        indexMap[num] = i
        result[i] = -1 // Initialize with -1
    }
    
    // Process nums2 and update result directly
    for i, num := range nums2 {
        for len(stack) > 0 && num > stack[len(stack)-1] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if idx, exists := indexMap[top]; exists {
                result[idx] = num
            }
        }
        stack = append(stack, num)
    }
    
    return result
}
```
- **Pros**: Single pass, no separate mapping phase
- **Cons**: More complex logic, requires index mapping

#### 2. Brute Force (O(N × M) time, O(1) space)
```go
func nextGreaterElementBrute(nums1 []int, nums2 []int) []int {
    result := make([]int, len(nums1))
    
    for i, num1 := range nums1 {
        found := false
        for j, num2 := range nums2 {
            if num2 == num1 {
                // Look for next greater
                for k := j + 1; k < len(nums2); k++ {
                    if nums2[k] > num2 {
                        result[i] = nums2[k]
                        found = true
                        break
                    }
                }
                if !found {
                    result[i] = -1
                }
                break
            }
        }
        if !found {
            result[i] = -1
        }
    }
    
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(N × M) time, too slow for large inputs

#### 3. Binary Search Approach (O(N log N + M log N) time, O(N) space)
```go
func nextGreaterElementBinary(nums1 []int, nums2 []int) []int {
    // Sort nums2 with original indices
    indexedNums2 := make([]struct{value int; index int}, len(nums2))
    for i, num := range nums2 {
        indexedNums2[i] = struct{value int; index int}{num, i}
    }
    sort.Slice(indexedNums2, func(i, j int) bool {
        return indexedNums2[i].value < indexedNums2[j].value
    })
    
    // For each nums1 element, find next greater using binary search
    result := make([]int, len(nums1))
    for i, num1 := range nums1 {
        // Binary search logic here
        result[i] = -1 // Simplified
    }
    
    return result
}
```
- **Pros**: Uses binary search efficiency
- **Cons**: Complex implementation, may not preserve order correctly

### Extensions for Interviews:
- **Circular Array**: Next Greater Element II problem
- **Next Smaller Element**: Find first smaller element to right
- **K-th Greater Element**: Find k-th greater element
- **Multiple Queries**: Handle multiple subsets efficiently
- **Streaming Data**: Process data as it arrives
*/
func main() {
	// Test cases
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
		{[]int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1}, []int{1}},
	}
	
	for i, tc := range testCases {
		result := nextGreaterElement(tc.nums1, tc.nums2)
		fmt.Printf("Test Case %d: nums1=%v, nums2=%v -> Next greater: %v\n", 
			i+1, tc.nums1, tc.nums2, result)
	}
}
