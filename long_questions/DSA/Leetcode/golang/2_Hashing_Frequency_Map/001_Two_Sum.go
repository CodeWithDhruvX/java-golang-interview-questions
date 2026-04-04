package main

import "fmt"

// 1. Two Sum
// Time: O(N), Space: O(N)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	
	return []int{}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hash Map for Complement Lookup
- **Hash Map**: Stores previously seen numbers and their indices
- **Complement Calculation**: For each number, calculate target - current
- **Constant Lookup**: Check if complement exists in hash map in O(1)
- **Single Pass**: Process array once, building map and checking simultaneously

## 2. PROBLEM CHARACTERISTICS
- **Pair Finding**: Need to find two numbers that sum to target
- **Index Return**: Must return indices, not the numbers themselves
- **Unique Solution**: Exactly one valid pair exists
- **No Reuse**: Cannot use the same element twice

## 3. SIMILAR PROBLEMS
- Two Sum II - Input array is sorted (LeetCode 167)
- Two Sum III - Data structure design (LeetCode 170)
- 3Sum, 4Sum - Extension to multiple numbers (LeetCode 15, 18)
- Subarray Sum Equals K (LeetCode 560)

## 4. KEY OBSERVATIONS
- **Complement Concept**: If a + b = target, then b = target - a
- **Hash Map Advantage**: O(1) lookup vs O(N) linear search
- **Order Independence**: Can find pair in either direction
- **Single Pass Efficiency**: Build map while checking for complements

## 5. VARIATIONS & EXTENSIONS
- **Sorted Input**: Can use two-pointer approach instead of hash
- **Multiple Solutions**: Return all pairs instead of just one
- **Data Structure**: Design a class supporting add/find operations
- **K-Sum Problem**: Find k numbers that sum to target

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is there exactly one solution? Can elements be reused?"
- Edge cases: empty array, single element, no solution
- Space complexity: O(N) for hash map storage
- Time complexity: O(N) - single pass through array

## 7. COMMON MISTAKES
- Using same element twice (need to check index difference)
- Not handling negative numbers correctly
- Returning numbers instead of indices
- Using nested loops instead of hash map (O(N²) solution)
- Forgetting to handle no solution case

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal** for unsorted arrays
- **For sorted arrays**: Use two-pointer approach with O(1) space
- **For multiple queries**: Pre-sort and use two-pointer or binary search
- **Memory optimization**: Use smaller integer types if possible

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a matching pair:**
- You're looking for two numbers that add up to a target
- As you scan through numbers, you remember what you've seen
- For each new number, you ask: "Did I already see its partner?"
- If yes, you found your pair! If no, you remember this number for future
- It's like having a perfect memory of all numbers you've encountered

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers and a target sum
2. **Goal**: Find indices of two numbers that sum to target
3. **Output**: Indices of the two numbers
4. **Constraint**: Exactly one solution exists

#### Phase 2: Key Insight Recognition
- **"Complement thinking"** → Instead of checking all pairs, find what we need
- **"Memory approach"** → Remember what we've seen to avoid re-scanning
- **"Hash map power"** → Instant lookup of previously seen numbers
- **"Single pass"** → Build memory while searching simultaneously

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find two numbers that add up to target.
For each number I see, I can calculate what I need: complement = target - current.
Instead of searching the whole array for this complement,
I'll keep a record (hash map) of all numbers I've seen so far.
When I encounter a new number, I'll check if its complement is already in my record.
If yes, I found my pair! If no, I'll add this number to my record for future checks."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty slice (no elements to pair)
- **Single element**: Return empty slice (need two elements)
- **No solution**: Problem guarantees exactly one solution
- **Duplicate elements**: Hash map handles correctly with indices

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums = [2, 7, 11, 15], target = 9

Human thinking:
"I'll scan through the array, remembering numbers I've seen.

Position 0: num = 2
   Complement needed: 9 - 2 = 7
   Have I seen 7 before? No (map is empty)
   Remember 2 at position 0: map = {2: 0}

Position 1: num = 7
   Complement needed: 9 - 7 = 2
   Have I seen 2 before? Yes! It's at position 0
   Found pair: [0, 1]
   Done! Return [0, 1]"
```

#### Phase 6: Intuition Validation
- **Why it works**: We check every possible pair exactly once
- **Why it's efficient**: Hash map gives O(1) lookup instead of O(N) search
- **Why it's correct**: We find the first valid pair, which is guaranteed unique

### Common Human Pitfalls & How to Avoid Them
1. **"Should I check all pairs?"** → No, use complement approach with hash map
2. **"What about the same element twice?"** → Check indices are different
3. **"Should I sort first?"** → Only if you need O(1) space and can lose indices
4. **"What if there are multiple solutions?"** → Return first one found

### Real-World Analogy
**Like finding a dance partner:**
- You have a list of people with their heights
- You need to find two people whose heights sum to a target
- As you meet each person, you remember their height
- When you meet someone new, you check if you already know their perfect partner
- If yes, you've found your dance pair!
- It's much faster than comparing everyone with everyone else

### Human-Readable Pseudocode
```
function twoSum(numbers, target):
    seenNumbers = empty map  // number -> index
    
    for each index, number in numbers:
        partnerNeeded = target - number
        
        if partnerNeeded exists in seenNumbers:
            return [seenNumbers[partnerNeeded], index]
        
        seenNumbers[number] = index
    
    return []  // no solution found
```

### Execution Visualization

### Example: nums = [2, 7, 11, 15], target = 9
```
Initial: numMap = {}

Step 1: i=0, num=2
→ complement = 9 - 2 = 7
→ 7 not in numMap
→ numMap[2] = 0
State: numMap = {2: 0}

Step 2: i=1, num=7
→ complement = 9 - 7 = 2
→ 2 exists in numMap at index 0!
→ Return [0, 1]

Final Answer: [0, 1]
```

### Example with more steps: nums = [3, 2, 4], target = 6
```
Initial: numMap = {}

Step 1: i=0, num=3
→ complement = 6 - 3 = 3
→ 3 not in numMap
→ numMap[3] = 0
State: numMap = {3: 0}

Step 2: i=1, num=2
→ complement = 6 - 2 = 4
→ 4 not in numMap
→ numMap[2] = 1
State: numMap = {3: 0, 2: 1}

Step 3: i=2, num=4
→ complement = 6 - 4 = 2
→ 2 exists in numMap at index 1!
→ Return [1, 2]

Final Answer: [1, 2]
```

### Key Visualization Points:
- **Complement calculation**: target - current number
- **Hash map lookup**: O(1) check for previously seen numbers
- **Index storage**: Store current number for future complement checks
- **Early termination**: Return as soon as pair is found

### Memory Layout Visualization:
```
Array:    [2][7][11][15]
Index:     0  1   2  3
Process:
i=0, num=2: complement=7, map={}
i=1, num=7: complement=2, map={2:0} → FOUND!
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Hash Map Operations**: O(1) average for insert and lookup
- **Overall**: O(N) time, O(N) space for hash map
- **Early Termination**: Often returns before processing entire array

### Alternative Approaches:

#### 1. Brute Force (O(N²))
```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}
```
- **Pros**: Simple, no extra space
- **Cons**: O(N²) time complexity

#### 2. Two Pointer (for sorted arrays, O(N))
```go
func twoSum(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return []int{left, right}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return []int{}
}
```
- **Pros**: O(N) time, O(1) space
- **Cons**: Requires sorted array, loses original indices

### Extensions for Interviews:
- **Multiple pairs**: Return all valid pairs
- **K-Sum problem**: Find k numbers that sum to target
- **Data structure design**: Support add/find operations
- **Streaming input**: Handle numbers as they arrive
*/
func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{-1, -2, -3, -4, -5}, -8},
	}
	
	for i, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Indices: %v\n", i+1, tc.nums, tc.target, result)
	}
}
