package main

import "fmt"

// 217. Contains Duplicate
// Time: O(N), Space: O(N)
func containsDuplicate(nums []int) bool {
	numSet := make(map[int]bool)
	
	for _, num := range nums {
		if _, exists := numSet[num]; exists {
			return true
		}
		numSet[num] = true
	}
	
	return false
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hash Set for Duplicate Detection
- **Hash Set Usage**: Track seen elements efficiently
- **Early Termination**: Return immediately when duplicate found
- **Constant Time Lookup**: O(1) average time for membership test
- **Space-Time Tradeoff**: Use extra space for O(N) time

## 2. PROBLEM CHARACTERISTICS
- **Duplicate Detection**: Check if any element appears more than once
- **Existence Query**: Only need to know if duplicate exists, not where
- **Early Exit**: Can stop as soon as first duplicate is found
- **Set Natural Fit**: Hash set perfect for tracking seen elements

## 3. SIMILAR PROBLEMS
- Contains Duplicate II (LeetCode 219) - Duplicates within distance k
- Contains Duplicate III (LeetCode 220) - Duplicates within distance and value range
- First Missing Positive (LeetCode 41) - Find smallest missing positive
- Find Duplicate Number (LeetCode 287) - Find duplicate in array

## 4. KEY OBSERVATIONS
- **Hash set ideal**: O(1) average lookup and insertion
- **Early termination**: Can stop when first duplicate found
- **Space requirement**: Need to store seen elements
- **Linear scan**: Single pass through array sufficient

## 5. VARIATIONS & EXTENSIONS
- **Distance constraint**: Duplicates within specific distance
- **Value range constraint**: Duplicates within specific value difference
- **Multiple duplicates**: Count all duplicates or find all duplicate pairs
- **Streaming data**: Handle data that arrives over time

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can I modify the input? What about space constraints?"
- Edge cases: empty array (false), single element (false), all same (true)
- Time complexity: O(N) time, O(N) space
- Alternative: O(N log N) time, O(1) space with sorting

## 7. COMMON MISTAKES
- Using nested loops (O(N²) time)
- Not handling empty array case
- Using array instead of hash set (inefficient lookups)
- Not returning early when duplicate found
- Forgetting negative numbers or zero

## 8. OPTIMIZATION STRATEGIES
- **Early termination**: Stop when first duplicate found
- **Pre-size hash set**: Estimate capacity to reduce reallocations
- **Bit set optimization**: For limited integer ranges
- **Sorting alternative**: If space is constrained

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like checking for duplicate tickets:**
- You're checking tickets (numbers) at an event entrance
- You want to know if anyone has a duplicate ticket
- You keep a list of tickets you've already seen
- For each new ticket, you check if it's already in your list
- If you find a duplicate, you can stop checking immediately

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: Determine if any element appears more than once
3. **Output**: Boolean (true if duplicate exists, false otherwise)
4. **Constraint**: Need efficient detection

#### Phase 2: Key Insight Recognition
- **"Hash set natural fit"** → Perfect for tracking seen elements
- **"Early termination"** → Can stop when first duplicate found
- **"Linear scan sufficient"** → No need for nested loops
- **"Space-time tradeoff"** → Use extra space for O(N) time

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if any number appears more than once.
I'll keep track of numbers I've seen using a hash set.
For each number, I'll check if it's already in the set.
If it is, I found a duplicate and can return true immediately.
If not, I'll add it to the set and continue.
If I finish the array without finding duplicates, return false."
```

#### Phase 4: Edge Case Handling
- **Empty array**: No duplicates possible, return false
- **Single element**: No duplicates possible, return false
- **All same elements**: First duplicate found at second element
- **Large arrays**: Handle efficiently with hash set

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1, 2, 3, 1]

Human thinking:
"I'll process each number and track what I've seen:

Number 1:
- Check if 1 is in seen set: No
- Add 1 to seen set: {1}

Number 2:
- Check if 2 is in seen set: No
- Add 2 to seen set: {1, 2}

Number 3:
- Check if 3 is in seen set: No
- Add 3 to seen set: {1, 2, 3}

Number 1:
- Check if 1 is in seen set: Yes!
- Found duplicate, return true immediately

No need to process further!"
```

#### Phase 6: Intuition Validation
- **Why hash set works**: O(1) average lookup and insertion
- **Why early termination**: Finding any duplicate is sufficient
- **Why O(N) time**: Single pass through array
- **Why O(N) space**: Need to store up to N elements

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use sorting?"** → Sorting is O(N log N) and modifies input
2. **"Should I use nested loops?"** → O(N²) time, too slow for large inputs
3. **"What about space constraints?"** → Sorting alternative if space limited
4. **"Can I optimize further?"** → Early termination is already optimal

### Real-World Analogy
**Like checking for duplicate IDs at a conference:**
- You're registering attendees and checking their ID numbers
- You want to ensure no two attendees have the same ID
- You keep a list of IDs you've already registered
- For each new attendee, you check if their ID is already in your list
- If you find a duplicate, you flag it immediately

### Human-Readable Pseudocode
```
function containsDuplicate(nums):
    seenSet = set()
    
    for num in nums:
        if num in seenSet:
            return true
        seenSet.add(num)
    
    return false
```

### Execution Visualization

### Example: [1, 2, 3, 1]
```
Hash Set Evolution:
Processing 1:
- 1 not in {} → add 1
- set = {1}

Processing 2:
- 2 not in {1} → add 2
- set = {1, 2}

Processing 3:
- 3 not in {1, 2} → add 3
- set = {1, 2, 3}

Processing 1:
- 1 is in {1, 2, 3} → FOUND DUPLICATE!
- Return true immediately

Result: true
```

### Key Visualization Points:
- **Hash set tracking**: Keep track of seen elements
- **Constant time lookup**: O(1) average time for membership test
- **Early termination**: Stop as soon as duplicate found
- **Linear progression**: Single pass through array

### Memory Layout Visualization:
```
Array Processing Flow:
[1, 2, 3, 1]
  ↓
{1}
  ↓
{1, 2}
  ↓
{1, 2, 3}
  ↓
{1, 2, 3} ← 1 already exists → DUPLICATE FOUND!

Hash Set Contents:
Step 1: {1}
Step 2: {1, 2}
Step 3: {1, 2, 3}
Step 4: {1, 2, 3} (duplicate detected)
```

### Time Complexity Breakdown:
- **Hash set approach**: O(N) time, O(N) space
- **Sorting approach**: O(N log N) time, O(1) space (if in-place)
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Bit set**: O(N) time, O(R) space where R is value range

### Alternative Approaches:

#### 1. Sorting Approach (O(N log N) time, O(1) space)
```go
func containsDuplicateSort(nums []int) bool {
    sort.Ints(nums)
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    
    return false
}
```
- **Pros**: O(1) extra space, no hash map overhead
- **Cons**: O(N log N) time, modifies input array

#### 2. Bit Set for Limited Range (O(N) time, O(R) space)
```go
func containsDuplicateBitSet(nums []int, maxValue int) bool {
    if maxValue > 1000000 { // Only for limited ranges
        return containsDuplicate(nums) // Fall back to hash set
    }
    
    bitSet := make([]uint64, (maxValue/64)+1)
    
    for _, num := range nums {
        if num < 0 {
            continue // Handle negative numbers separately
        }
        
        index := num / 64
        bit := uint64(1) << (num % 64)
        
        if bitSet[index]&bit != 0 {
            return true
        }
        bitSet[index] |= bit
    }
    
    return false
}
```
- **Pros**: Very memory efficient for limited ranges
- **Cons**: Only works for non-negative integers in limited range

#### 3. Floyd's Tortoise and Hare (for specific constraints)
```go
// Only works if numbers are in range [1, n-1] and array length is n
func containsDuplicateFloyd(nums []int) bool {
    // This is for Find Duplicate Number problem, not general case
    // Included here for educational purposes
    return false // Not applicable for general duplicate detection
}
```
- **Pros**: O(1) space for specific constraints
- **Cons**: Only works for very specific problem variants

### Extensions for Interviews:
- **Contains Duplicate II**: Duplicates within distance k
- **Contains Duplicate III**: Duplicates within distance and value range
- **Find All Duplicates**: Return all elements that appear more than once
- **Streaming Data**: Handle data that arrives over time
- **Count Duplicates**: Count how many duplicates exist
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 2, 2, 2},
		{},
		{0},
		{-1, -2, -3, -1},
	}
	
	for i, nums := range testCases {
		result := containsDuplicate(nums)
		fmt.Printf("Test Case %d: %v -> Contains duplicate: %t\n", i+1, nums, result)
	}
}
