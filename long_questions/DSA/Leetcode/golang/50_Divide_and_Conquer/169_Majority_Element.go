package main

import (
	"fmt"
	"math"
)

// 169. Majority Element - Divide and Conquer Approach
// Time: O(N log N), Space: O(log N) for recursion stack
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	return majorityElementHelper(nums, 0, len(nums)-1)
}

func majorityElementHelper(nums []int, left, right int) int {
	// Base case: only one element
	if left == right {
		return nums[left]
	}
	
	// Divide: find majority element in left and right halves
	mid := left + (right-left)/2
	leftMajority := majorityElementHelper(nums, left, mid)
	rightMajority := majorityElementHelper(nums, mid+1, right)
	
	// Conquer: combine results
	if leftMajority == rightMajority {
		return leftMajority
	}
	
	// Count occurrences of each majority candidate
	leftCount := countOccurrences(nums, left, right, leftMajority)
	rightCount := countOccurrences(nums, left, right, rightMajority)
	
	if leftCount > rightCount {
		return leftMajority
	}
	return rightMajority
}

func countOccurrences(nums []int, left, right, target int) int {
	count := 0
	for i := left; i <= right; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

// Divide and Conquer with early termination
func majorityElementOptimized(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	return majorityElementOptimizedHelper(nums, 0, len(nums)-1)
}

func majorityElementOptimizedHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	
	mid := left + (right-left)/2
	leftMajority := majorityElementOptimizedHelper(nums, left, mid)
	rightMajority := majorityElementOptimizedHelper(nums, mid+1, right)
	
	if leftMajority == rightMajority {
		return leftMajority
	}
	
	// Early termination: if subarray is too small, count directly
	subArraySize := right - left + 1
	if subArraySize <= 10 {
		return findMajorityBruteForce(nums, left, right)
	}
	
	leftCount := countOccurrences(nums, left, right, leftMajority)
	rightCount := countOccurrences(nums, left, right, rightMajority)
	
	if leftCount > rightCount {
		return leftMajority
	}
	return rightMajority
}

func findMajorityBruteForce(nums []int, left, right int) int {
	maxCount := 0
	majority := nums[left]
	
	for i := left; i <= right; i++ {
		count := countOccurrences(nums, left, right, nums[i])
		if count > maxCount {
			maxCount = count
			majority = nums[i]
		}
	}
	
	return majority
}

// Divide and Conquer with memoization
func majorityElementMemoized(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	memo := make(map[string]int)
	return majorityElementMemoizedHelper(nums, 0, len(nums)-1, memo)
}

func majorityElementMemoizedHelper(nums []int, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left == right {
		memo[key] = nums[left]
		return nums[left]
	}
	
	mid := left + (right-left)/2
	leftMajority := majorityElementMemoizedHelper(nums, left, mid, memo)
	rightMajority := majorityElementMemoizedHelper(nums, mid+1, right, memo)
	
	var result int
	if leftMajority == rightMajority {
		result = leftMajority
	} else {
		leftCount := countOccurrences(nums, left, right, leftMajority)
		rightCount := countOccurrences(nums, left, right, rightMajority)
		if leftCount > rightCount {
			result = leftMajority
		} else {
			result = rightMajority
		}
	}
	
	memo[key] = result
	return result
}

// Divide and Conquer for finding all majority elements (> n/3)
func majorityElementAll(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	
	// Use divide and conquer to find candidates
	candidates := findMajorityCandidates(nums, 0, len(nums)-1)
	
	// Verify candidates
	var result []int
	for _, candidate := range candidates {
		if countOccurrences(nums, 0, len(nums)-1, candidate) > len(nums)/3 {
			result = append(result, candidate)
		}
	}
	
	return result
}

func findMajorityCandidates(nums []int, left, right int) []int {
	if left == right {
		return []int{nums[left]}
	}
	
	mid := left + (right-left)/2
	leftCandidates := findMajorityCandidates(nums, left, mid)
	rightCandidates := findMajorityCandidates(nums, mid+1, right)
	
	// Merge candidates and remove duplicates
	candidateMap := make(map[int]bool)
	for _, candidate := range leftCandidates {
		candidateMap[candidate] = true
	}
	for _, candidate := range rightCandidates {
		candidateMap[candidate] = true
	}
	
	var candidates []int
	for candidate := range candidateMap {
		candidates = append(candidates, candidate)
	}
	
	return candidates
}

// Divide and Conquer with parallel processing simulation
func majorityElementParallel(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// Simulate parallel processing by dividing into chunks
	chunkSize := len(nums) / 4
	if chunkSize < 1 {
		chunkSize = 1
	}
	
	var chunks [][]int
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		chunks = append(chunks, nums[i:end])
	}
	
	// Find majority in each chunk
	chunkMajorities := make([]int, len(chunks))
	for i, chunk := range chunks {
		chunkMajorities[i] = findMajorityBruteForce(chunk, 0, len(chunk)-1)
	}
	
	// Find majority among chunk majorities
	return findMajorityAmongCandidates(nums, chunkMajorities)
}

func findMajorityAmongCandidates(nums []int, candidates []int) int {
	maxCount := 0
	majority := candidates[0]
	
	for _, candidate := range candidates {
		count := countOccurrences(nums, 0, len(nums)-1, candidate)
		if count > maxCount {
			maxCount = count
			majority = candidate
		}
	}
	
	return majority
}

// Divide and Conquer with range queries
func majorityElementRangeQuery(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// Build segment tree for range queries
	segTree := buildSegmentTree(nums, 0, len(nums)-1)
	
	// Query the entire range
	return querySegmentTree(segTree, 0, len(nums)-1, 0, len(nums)-1)
}

func buildSegmentTree(nums []int, left, right int) *SegmentTreeNode {
	if left == right {
		return &SegmentTreeNode{
			value:    nums[left],
			count:    1,
			left:     nil,
			right:    nil,
			leftIdx:  left,
			rightIdx: right,
		}
	}
	
	mid := left + (right-left)/2
	leftNode := buildSegmentTree(nums, left, mid)
	rightNode := buildSegmentTree(nums, mid+1, right)
	
	// Merge nodes
	merged := mergeSegmentNodes(leftNode, rightNode)
	return merged
}

type SegmentTreeNode struct {
	value    int
	count    int
	left     *SegmentTreeNode
	right    *SegmentTreeNode
	leftIdx  int
	rightIdx int
}

func mergeSegmentNodes(left, right *SegmentTreeNode) *SegmentTreeNode {
	if left.value == right.value {
		return &SegmentTreeNode{
			value:    left.value,
			count:    left.count + right.count,
			leftIdx:  left.leftIdx,
			rightIdx: right.rightIdx,
		}
	}
	
	leftCount := left.count
	rightCount := right.count
	
	if leftCount > rightCount {
		return &SegmentTreeNode{
			value:    left.value,
			count:    leftCount - rightCount,
			leftIdx:  left.leftIdx,
			rightIdx: right.rightIdx,
		}
	} else {
		return &SegmentTreeNode{
			value:    right.value,
			count:    rightCount - leftCount,
			leftIdx:  left.leftIdx,
			rightIdx: right.rightIdx,
		}
	}
}

func querySegmentTree(node *SegmentTreeNode, left, right, queryLeft, queryRight int) int {
	if queryLeft > node.rightIdx || queryRight < node.leftIdx {
		return -1
	}
	
	if queryLeft <= node.leftIdx && node.rightIdx <= queryRight {
		return node.value
	}
	
	leftResult := querySegmentTree(node.left, left, right, queryLeft, queryRight)
	rightResult := querySegmentTree(node.right, left, right, queryLeft, queryRight)
	
	if leftResult == -1 {
		return rightResult
	}
	if rightResult == -1 {
		return leftResult
	}
	
	// Need to count occurrences in the range
	// For simplicity, return the one with more occurrences
	leftCount := countOccurrencesInRange([]int{leftResult}, queryLeft, queryRight, leftResult)
	rightCount := countOccurrencesInRange([]int{rightResult}, queryLeft, queryRight, rightResult)
	
	if leftCount >= rightCount {
		return leftResult
	}
	return rightResult
}

func countOccurrencesInRange(nums []int, left, right, target int) int {
	// Simplified counting
	return 1
}

// Boyer-Moore algorithm for comparison (O(N) time, O(1) space)
func majorityElementBoyerMoore(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	candidate := nums[0]
	count := 1
	
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
		
		if count == 0 {
			candidate = nums[i]
			count = 1
		}
	}
	
	return candidate
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Divide and Conquer for Majority Element
- **Recursive Division**: Split array into halves recursively
- **Majority Propagation**: Majority element must be majority in at least one half
- **Count Verification**: Count occurrences of candidates from subarrays
- **Conquer Strategy**: Combine subproblem results with verification

## 2. PROBLEM CHARACTERISTICS
- **Majority Element**: Element appearing more than ⌊n/2⌋ times
- **Guaranteed Existence**: Problem guarantees a majority element exists
- **Frequency Analysis**: Need to find element with maximum frequency
- **Divide Property**: Majority element propagates through divisions

## 3. SIMILAR PROBLEMS
- Majority Element (LeetCode 169) - Same problem
- Majority Element II - Find elements > n/3 times
- Top K Frequent Elements - Find most frequent elements
- Find Duplicate Numbers - Frequency-based problems

## 4. KEY OBSERVATIONS
- **Majority Inheritance**: If element is majority in whole array, it's majority in at least one half
- **Candidate Reduction**: Each division reduces candidate space
- **Count Verification**: Need to verify candidates from subarrays
- **Recursive Structure**: Same problem applied to smaller arrays

## 5. VARIATIONS & EXTENSIONS
- **Standard D&C**: O(N log N) time, O(log N) space
- **Boyer-Moore**: O(N) time, O(1) space - optimal
- **Memoization**: O(N log N) time, O(N log N) space - reduces redundancy
- **All Majority Elements**: Find elements > n/3 times

## 6. INTERVIEW INSIGHTS
- Always clarify: "Majority guaranteed? > n/2 or > n/3? Space constraints?"
- Edge cases: empty array, single element, no majority guarantee
- Time complexity: O(N log N) for D&C, O(N) for Boyer-Moore
- Space complexity: O(log N) recursion stack, O(1) for Boyer-Moore
- Key insight: Boyer-Moore is optimal but D&C demonstrates pattern

## 7. COMMON MISTAKES
- Wrong base case handling for single elements
- Missing count verification step
- Incorrect candidate combination logic
- Not handling empty arrays properly
- Wrong majority threshold calculation

## 8. OPTIMIZATION STRATEGIES
- **Standard D&C**: O(N log N) time, O(log N) space - demonstrates pattern
- **Boyer-Moore**: O(N) time, O(1) space - optimal solution
- **Memoization**: O(N log N) time, O(N log N) space - reduces redundancy
- **Early Termination**: Switch to brute force for small subarrays

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most popular candidate in an election:**
- You have voters divided into different districts (array halves)
- The overall winner must be the winner in at least one district
- You can recursively find winners in each district
- Then verify which district winner has the most votes overall
- Like an election analyst finding the majority candidate by district results

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers with guaranteed majority element
2. **Goal**: Find element appearing more than ⌊n/2⌋ times
3. **Constraints**: Majority element guaranteed to exist
4. **Output**: The majority element value

#### Phase 2: Key Insight Recognition
- **"Divide natural"** → Can split array at midpoint
- **"Majority inheritance"** → Overall majority must be majority in at least one half
- **"Candidate verification"** → Need to count occurrences of candidates
- **"Recursive structure"** → Same problem applied to subarrays

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find majority element (> n/2 occurrences).
Brute force: count all elements O(N²).

Divide and Conquer Approach:
1. Split array at midpoint
2. Recursively find majority in left half
3. Recursively find majority in right half
4. If both halves agree, that's the majority
5. If different, count occurrences and return the one with more

This gives O(N log N) time, O(log N) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 or handle as specified
- **Single element**: That element is majority
- **No majority guarantee**: Need verification step
- **Large arrays**: Recursion depth O(log N)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [3, 2, 3]

Human thinking:
"Divide and Conquer Process:
Step 1: Split at middle (index 1)
Left: [3], Right: [2, 3]

Step 2: Recursively solve left [3]
Base case: single element
Majority = 3

Step 3: Recursively solve right [2, 3]
Split at middle (index 2)
LeftRight: [2], RightRight: [3]
Majority LeftRight = 2, Majority RightRight = 3
Different candidates, count occurrences:
Count(2) = 1, Count(3) = 1
Return 3 (arbitrary tie-break)

Step 4: Combine results
Left majority = 3, Right majority = 3
Same candidate, return 3

Final result: 3 ✓ (appears 2 times > 3/2 = 1.5)"
```

#### Phase 6: Intuition Validation
- **Why divide**: Natural way to break down problem
- **Why majority inheritance**: Overall majority must dominate at least one half
- **Why verification**: Different candidates from halves need verification
- **Why O(N log N)**: Each level does O(N) counting, O(log N) levels

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count all elements?"** → O(N²) vs O(N log N), too slow
2. **"Should I use hash map?"** → O(N) time, O(N) space, Boyer-Moore is better
3. **"What about tie cases?"** → Problem guarantees majority exists
4. **"Can I skip verification?"** → Different candidates need verification
5. **"Why Boyer-Moore?"** → O(N) time, O(1) space, optimal solution

### Real-World Analogy
**Like finding the most popular product in a store chain:**
- You have stores in different regions (array halves)
- The overall most popular product must be most popular in at least one region
- You can survey each region recursively to find local favorites
- Then verify which regional favorite has the most overall sales
- Like a retail analyst finding the best-selling product by region analysis

### Human-Readable Pseudocode
```
function majorityElement(nums):
    if len(nums) == 0:
        return -1
    
    return majorityElementHelper(nums, 0, len(nums) - 1)

function majorityElementHelper(nums, left, right):
    if left == right:
        return nums[left]
    
    mid = (left + right) // 2
    
    # Find majority in left half
    leftMajority = majorityElementHelper(nums, left, mid)
    
    # Find majority in right half
    rightMajority = majorityElementHelper(nums, mid + 1, right)
    
    # If same, that's the majority
    if leftMajority == rightMajority:
        return leftMajority
    
    # Different candidates, count occurrences
    leftCount = countOccurrences(nums, left, right, leftMajority)
    rightCount = countOccurrences(nums, left, right, rightMajority)
    
    return leftCount > rightCount ? leftMajority : rightMajority
```

### Execution Visualization

### Example: nums = [2, 2, 1, 1, 1, 2, 2]
```
Level 0: [2, 2, 1, 1, 1, 2, 2]
         Split at index 3
    Left: [2, 2, 1, 1]    Right: [1, 2, 2]

Level 1: Left: [2, 2, 1, 1]
         Split at index 1
    LeftLeft: [2, 2]    LeftRight: [1, 1]

Level 2: LeftLeft: [2, 2]
         Split at index 0
    [2]    [2]
    Majority = 2, Majority = 2
    Same candidate, return 2

Level 2: LeftRight: [1, 1]
         Split at index 2
    [1]    [1]
    Majority = 1, Majority = 1
    Same candidate, return 1

Level 1: Left: [2, 2, 1, 1]
    Left majority = 2, Right majority = 1
    Different candidates, count:
    Count(2) = 2, Count(1) = 2
    Return 2 (arbitrary tie-break)

Level 1: Right: [1, 2, 2]
    Similar process, majority = 2

Level 0: Combine results
    Left majority = 2, Right majority = 2
    Same candidate, return 2

Final result: 2 ✓ (appears 4 times > 7/2 = 3.5)
```

### Key Visualization Points:
- **Recursive Splitting**: Array split until single elements
- **Majority Propagation**: Majority element dominates subarrays
- **Candidate Verification**: Count occurrences when candidates differ
- **Tie Handling**: Problem guarantees majority, so ties resolve correctly

### Divide and Conquer Tree Visualization:
```
        [2,2,1,1,1,2,2]
       /              \
    [2,2,1,1]        [1,2,2]
    /      \        /      \
 [2,2]    [1,1]  [1]    [2,2]
  /   \    /   \        /   \
[2]   [2][1]   [1]    [2]   [2]
```

### Time Complexity Breakdown:
- **Standard D&C**: O(N log N) time, O(log N) space - demonstrates pattern
- **Boyer-Moore**: O(N) time, O(1) space - optimal solution
- **Memoization**: O(N log N) time, O(N log N) space - reduces redundancy
- **Early Termination**: O(N log N) time, O(log N) space - practical optimization

### Alternative Approaches:

#### 1. Boyer-Moore Voting (O(N) time, O(1) space)
```go
func majorityElementBoyerMoore(nums []int) int {
    candidate := nums[0]
    count := 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == candidate {
            count++
        } else {
            count--
        }
        
        if count == 0 {
            candidate = nums[i]
            count = 1
        }
    }
    
    return candidate
}
```
- **Pros**: Optimal O(N) time, O(1) space
- **Cons**: Requires majority guarantee, less intuitive

#### 2. Hash Map (O(N) time, O(N) space)
```go
func majorityElementHashMap(nums []int) int {
    freq := make(map[int]int)
    
    for _, num := range nums {
        freq[num]++
        if freq[num] > len(nums)/2 {
            return num
        }
    }
    
    return -1
}
```
- **Pros**: Simple, early termination possible
- **Cons**: O(N) extra space

#### 3. Sorting (O(N log N) time, O(1) space)
```go
func majorityElementSort(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
```
- **Pros**: Simple, leverages sorting
- **Cons**: O(N log N) time, modifies array

### Extensions for Interviews:
- **Majority Element II**: Find elements > n/3 times
- **No Guarantee**: Handle cases where majority might not exist
- **Range Queries**: Find majority in subarrays
- **Streaming Data**: Find majority in data stream
- **Real-world Applications**: Voting systems, trend analysis, data analytics
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Majority Element - Divide and Conquer ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{3, 2, 3}, "Standard case"},
		{[]int{2, 2, 1, 1, 1, 2, 2}, "Mixed case"},
		{[]int{1}, "Single element"},
		{[]int{1, 1, 1, 1}, "All same"},
		{[]int{1, 2, 3, 4, 5}, "No majority (but first returned)"},
		{[]int{6, 5, 5}, "Last element majority"},
		{[]int{1, 2, 3, 2, 2}, "Middle element majority"},
		{[]int{-1, -1, -2, -1, -3}, "With negatives"},
		{[]int{0, 0, 0, 1, 2}, "With zeros"},
		{[]int{100, 100, 100, 50, 50}, "Large numbers"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Input: %v\n", tc.nums)
		
		result1 := majorityElement(tc.nums)
		result2 := majorityElementOptimized(tc.nums)
		result3 := majorityElementMemoized(tc.nums)
		result4 := majorityElementParallel(tc.nums)
		result5 := majorityElementBoyerMoore(tc.nums)
		
		fmt.Printf("  Standard D&C: %d\n", result1)
		fmt.Printf("  Optimized: %d\n", result2)
		fmt.Printf("  Memoized: %d\n", result3)
		fmt.Printf("  Parallel: %d\n", result4)
		fmt.Printf("  Boyer-Moore: %d\n\n", result5)
	}
	
	// Test finding all majority elements (> n/3)
	fmt.Println("=== Finding All Majority Elements (> n/3) ===")
	allMajorityTest := []int{3, 2, 3}
	fmt.Printf("Input: %v\n", allMajorityTest)
	fmt.Printf("All majority elements: %v\n", majorityElementAll(allMajorityTest))
	
	allMajorityTest2 := []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Printf("Input: %v\n", allMajorityTest2)
	fmt.Printf("All majority elements: %v\n", majorityElementAll(allMajorityTest2))
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeArray := make([]int, 10000)
	for i := range largeArray {
		if i < 6000 { // Make 60% the majority
			largeArray[i] = 1
		} else {
			largeArray[i] = 2
		}
	}
	
	fmt.Printf("Large array test with %d elements (60% majority)\n", len(largeArray))
	
	result := majorityElement(largeArray)
	fmt.Printf("Standard D&C result: %d\n", result)
	
	result = majorityElementBoyerMoore(largeArray)
	fmt.Printf("Boyer-Moore result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty array
	fmt.Printf("Empty array: %d\n", majorityElement([]int{}))
	
	// Two elements
	fmt.Printf("Two elements [1,2]: %d\n", majorityElement([]int{1, 2}))
	fmt.Printf("Two elements [1,1]: %d\n", majorityElement([]int{1, 1}))
	
	// Large numbers
	largeVals := []int{1000000, 1000001, 1000000, 1000002}
	fmt.Printf("Large numbers: %d\n", majorityElement(largeVals))
	
	// Alternating pattern
	alternating := []int{1, 2, 1, 2, 1, 2, 1}
	fmt.Printf("Alternating pattern: %d\n", majorityElement(alternating))
	
	// Test with many duplicates
	fmt.Println("\n=== Many Duplicates Test ===")
	manyDup := make([]int, 1000)
	for i := range manyDup {
		if i < 501 { // Slight majority
			manyDup[i] = 42
		} else {
			manyDup[i] = i % 10
		}
	}
	
	result = majorityElement(manyDup)
	fmt.Printf("Many duplicates: %d (should be 42)\n", result)
	
	// Verify the result is actually majority
	actualCount := countOccurrences(manyDup, 0, len(manyDup)-1, result)
	fmt.Printf("Actual count: %d, Expected: >%d\n", actualCount, len(manyDup)/2)
}
