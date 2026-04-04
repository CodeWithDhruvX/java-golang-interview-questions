package main

import "fmt"

// 287. Find the Duplicate Number
// Time: O(N), Space: O(1) - Floyd's Tortoise and Hare Algorithm
func findDuplicate(nums []int) int {
	// Phase 1: Find the intersection point
	slow := nums[0]
	fast := nums[0]
	
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	// Phase 2: Find the entrance to the cycle
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}

// Alternative solution using cyclic sort (modifies the array)
func findDuplicateCyclicSort(nums []int) int {
	i := 0
	n := len(nums)
	
	for i < n {
		correctPos := nums[i] - 1
		if nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// The duplicate will be at the position where the number doesn't match index+1
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return nums[i]
		}
	}
	
	return -1 // Should never reach here for valid input
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Cycle Detection in Arrays
- **Floyd's Algorithm**: Tortoise and Hare cycle detection
- **Linked List Cycle**: Array indices form linked list with cycle
- **Cycle Entry Point**: Duplicate number is cycle entrance
- **Cyclic Sort Alternative**: In-place sorting to find duplicate

## 2. PROBLEM CHARACTERISTICS
- **Single Duplicate**: Exactly one number appears twice
- **Range Constraint**: Numbers from 1 to n in array of size n+1
- **No Modification**: Cannot modify array (Floyd's algorithm)
- **Space Efficiency**: O(1) space requirement

## 3. SIMILAR PROBLEMS
- Find the Duplicate Number (LeetCode 287) - Same problem
- Linked List Cycle (LeetCode 142) - Same Floyd algorithm
- Happy Number (LeetCode 202) - Cycle detection variant
- Circular Array Loop (LeetCode 457) - Cycle detection in arrays

## 4. KEY OBSERVATIONS
- **Linked List Mapping**: nums[i] acts as next pointer
- **Cycle Formation**: Duplicate creates cycle in linked list
- **Floyd's Algorithm**: Detect and find cycle entrance
- **Mathematical Guarantee**: Pigeonhole principle ensures duplicate exists

## 5. VARIATIONS & EXTENSIONS
- **Floyd's Algorithm**: Non-modifying, elegant solution
- **Cyclic Sort**: Modifying array, intuitive approach
- **Binary Search**: Count-based approach
- **Hash Set**: Simple but uses extra space

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can I modify array? Space constraints? Multiple duplicates?"
- Edge cases: adjacent duplicates, all same elements
- Time complexity: O(N) for Floyd's, O(N) for cyclic sort
- Space complexity: O(1) for both approaches
- Key insight: treat array as linked list with cycle

## 7. COMMON MISTAKES
- Wrong Floyd's algorithm implementation
- Not handling array bounds properly
- Infinite loop in cyclic sort with duplicates
- Missing edge cases for single element
- Wrong index calculations

## 8. OPTIMIZATION STRATEGIES
- **Floyd's Algorithm**: O(N) time, O(1) space - optimal, non-modifying
- **Cyclic Sort**: O(N) time, O(1) space - modifies array
- **Binary Search**: O(N log N) time, O(1) space - counting approach
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a meeting point in a maze:**
- Each array element points to the next position
- The duplicate creates a loop (people keep meeting at same spot)
- One person walks slowly (tortoise), one runs fast (hare)
- The fast person will eventually lap the slow person
- The meeting point reveals there's a cycle
- Like two people running on a circular track

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of n+1 integers from 1 to n (one duplicate)
2. **Goal**: Find the duplicate number
3. **Constraints**: O(1) space, cannot modify array
4. **Output**: The duplicate integer

#### Phase 2: Key Insight Recognition
- **"Linked list natural"** → nums[i] as next pointer forms linked list
- **"Cycle guaranteed"** → Duplicate creates cycle in linked list
- **"Floyd's algorithm"** → Classic solution for cycle detection
- **"Cycle entrance"** → Duplicate is where cycle starts

#### Phase 3: Strategy Development
```
Human thought process:
"I need duplicate without extra space.
Brute force: use hash set O(N) space.

Floyd's Algorithm Approach:
1. Treat array as linked list: next = nums[current]
2. Phase 1: Find intersection of slow and fast pointers
3. Phase 2: Find cycle entrance (duplicate)
4. Slow pointer moves 1 step, fast moves 2 steps
5. When they meet, reset slow to start
6. Move both 1 step until they meet again

This gives O(N) time, O(1) space!"
```

#### Phase 4: Edge Case Handling
- **Single duplicate**: Both algorithms work correctly
- **Adjacent duplicates**: Floyd's algorithm handles all cases
- **All same elements**: Cycle starts at first element
- **Large arrays**: Algorithm scales linearly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [1, 3, 4, 2, 2]

Human thinking:
"Floyd's Algorithm Process:
Phase 1: Find intersection
Step 1: slow=nums[0]=1, fast=nums[0]=1
Step 2: slow=nums[1]=3, fast=nums[nums[1]]=nums[3]=2
Step 3: slow=nums[3]=2, fast=nums[nums[2]]=nums[4]=2
They meet at index 3 (value 2) ✓

Phase 2: Find cycle entrance
Step 1: Reset slow to start: slow=nums[0]=1, fast stays at 2
Step 2: slow=nums[1]=3, fast=nums[2]=4
Step 3: slow=nums[3]=2, fast=nums[4]=2
They meet at index 3 (value 2) ✓

Result: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why linked list**: nums[i] naturally forms next pointer relationship
- **Why cycle forms**: Duplicate points to same position
- **Why Floyd's works**: Mathematical guarantee of meeting
- **Why cycle entrance**: First repeated element is duplicate

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N), may modify array
2. **"Should I use hash set?"** → Yes, but violates O(1) space
3. **"What about binary search?"** → Yes, counting approach works
4. **"Can I modify array?"** → Cyclic sort works if allowed
5. **"Why two phases?"** → First finds cycle, second finds entrance

### Real-World Analogy
**Like finding a popular meeting spot in an office:**
- Each office has instructions to go to another office
- One popular office is mentioned twice (duplicate)
- People following instructions eventually form a loop
- Two people start, one walks slowly, one runs fast
- The fast person eventually catches up to slow person
- The meeting point reveals the popular office
- Like detectives finding a frequently visited location

### Human-Readable Pseudocode
```
function findDuplicate(nums):
    # Phase 1: Find intersection
    slow = nums[0]
    fast = nums[0]
    
    while True:
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast:
            break
    
    # Phase 2: Find cycle entrance
    slow = nums[0]
    while slow != fast:
        slow = nums[slow]
        fast = nums[fast]
    
    return slow

# Alternative: Cyclic Sort (modifies array)
function findDuplicateCyclicSort(nums):
    i = 0
    while i < len(nums):
        correctPos = nums[i] - 1
        if nums[i] != nums[correctPos]:
            swap(nums[i], nums[correctPos])
        else:
            i += 1
    
    # Find first position with wrong number
    for i from 0 to len(nums)-1:
        if nums[i] != i + 1:
            return nums[i]
```

### Execution Visualization

### Example: nums = [1, 3, 4, 2, 2]
```
Floyd's Algorithm Process:

Array as Linked List:
0 → 1 → 3 → 2 → 4 → 2 → 4 → 2 → ...
Indices:  0    1    3    2    4    2    4
Values:  1    3    2    4    2    4    2

Phase 1: Find intersection
Step 1: slow=1, fast=1
Step 2: slow=3, fast=2
Step 3: slow=2, fast=2 (meet at index 3)

Phase 2: Find cycle entrance
Step 1: slow=1 (reset), fast=2
Step 2: slow=3, fast=4
Step 3: slow=2, fast=2 (meet at index 3)

Cycle entrance: value 2 at index 3
Result: 2 ✓
```

### Key Visualization Points:
- **Linked List Mapping**: nums[i] as next pointer
- **Cycle Formation**: Duplicate creates loop
- **Two-Phase Process**: Find cycle, then entrance
- **Meeting Point**: Reveals cycle structure

### Cycle Detection Visualization:
```
Array: [1, 3, 4, 2, 2]
Indices: 0→1→3→2→4→2→4→2→...
Values:  1→3→2→4→2→4→2→4→...

Cycle: 2→4→2 (duplicate 2 creates cycle)
Entrance: First 2 in the cycle
```

### Time Complexity Breakdown:
- **Floyd's Algorithm**: O(N) time, O(1) space - optimal, non-modifying
- **Cyclic Sort**: O(N) time, O(1) space - modifies array
- **Binary Search**: O(N log N) time, O(1) space - counting approach
- **Hash Set**: O(N) time, O(N) space - simple but uses extra space

### Alternative Approaches:

#### 1. Binary Search (O(N log N) time, O(1) space)
```go
func findDuplicateBinarySearch(nums []int) int {
    left, right := 1, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        count := 0
        
        for _, num := range nums {
            if num <= mid {
                count++
            }
        }
        
        if count > mid {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}
```
- **Pros**: O(1) space, no modification
- **Cons**: O(N log N) time, less intuitive

#### 2. Hash Set (O(N) time, O(N) space)
```go
func findDuplicateHash(nums []int) int {
    seen := make(map[int]bool)
    
    for _, num := range nums {
        if seen[num] {
            return num
        }
        seen[num] = true
    }
    
    return -1
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Uses O(N) extra space

#### 3. Sum Method (O(N) time, O(1) space)
```go
func findDuplicateSum(nums []int) int {
    n := len(nums) - 1
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    
    for _, num := range nums {
        actualSum += num
    }
    
    return actualSum - expectedSum
}
```
- **Pros**: Simple arithmetic, O(1) space
- **Cons**: Only works for exactly one duplicate

### Extensions for Interviews:
- **Multiple Duplicates**: Find all repeated numbers
- **K Duplicates**: Find numbers appearing k times
- **Distance to Duplicate**: Find distance to duplicate
- **Cycle Length**: Find length of the cycle
- **Real-world Applications**: Memory leak detection, loop detection
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
		{1, 1},
		{2, 2, 2, 2, 2},
		{1, 4, 4, 3, 2},
		{5, 4, 3, 2, 1, 5},
		{3, 1, 2, 3, 4, 5},
		{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy for Floyd's algorithm (doesn't modify array)
		numsCopy1 := make([]int, len(nums))
		copy(numsCopy1, nums)
		
		// Make a copy for cyclic sort (modifies array)
		numsCopy2 := make([]int, len(nums))
		copy(numsCopy2, nums)
		
		result1 := findDuplicate(numsCopy1)
		result2 := findDuplicateCyclicSort(numsCopy2)
		
		fmt.Printf("Test Case %d: %v -> Duplicate (Floyd): %d, Duplicate (Cyclic): %d\n", 
			i+1, nums, result1, result2)
	}
}
