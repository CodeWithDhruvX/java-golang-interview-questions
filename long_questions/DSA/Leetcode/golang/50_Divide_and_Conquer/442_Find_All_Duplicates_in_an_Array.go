package main

import (
	"fmt"
	"math"
)

// 442. Find All Duplicates in an Array - Divide and Conquer Approach
// Time: O(N log N), Space: O(log N) for recursion stack
func findDuplicates(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Sort the array first (divide and conquer)
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	mergeSort(sorted, 0, len(sorted)-1)
	
	// Find duplicates in sorted array
	var duplicates []int
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			duplicates = append(duplicates, sorted[i])
			// Skip all additional occurrences
			for i+1 < len(sorted) && sorted[i+1] == sorted[i] {
				i++
			}
		}
	}
	
	return duplicates
}

func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	
	// Merge two sorted halves
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	
	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}
	
	// Copy back to original array
	for i, val := range temp {
		nums[left+i] = val
	}
}

// Divide and conquer with counting sort approach
func findDuplicatesCountingSort(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Find min and max to determine range
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	
	// Count frequency of each number
	rangeSize := maxVal - minVal + 1
	count := make([]int, rangeSize)
	
	for _, num := range nums {
		count[num-minVal]++
	}
	
	// Find duplicates
	var duplicates []int
	for i, freq := range count {
		if freq > 1 {
			duplicates = append(duplicates, i+minVal)
		}
	}
	
	return duplicates
}

// Divide and conquer with hash map (more efficient)
func findDuplicatesHashMap(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Use divide and conquer to build frequency map
	freq := make(map[int]int)
	buildFrequencyMap(nums, 0, len(nums)-1, freq)
	
	// Extract duplicates
	var duplicates []int
	for num, count := range freq {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}
	
	return duplicates
}

func buildFrequencyMap(nums []int, left, right int, freq map[int]int) {
	if left > right {
		return
	}
	
	if left == right {
		freq[nums[left]]++
		return
	}
	
	mid := left + (right-left)/2
	buildFrequencyMap(nums, left, mid, freq)
	buildFrequencyMap(nums, mid+1, right, freq)
}

// Divide and conquer with binary search tree
func findDuplicatesBST(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Build BST and count frequencies
	root := nil
	freq := make(map[int]int)
	
	for _, num := range nums {
		root = insertBST(root, num, freq)
	}
	
	// Extract duplicates
	var duplicates []int
	for num, count := range freq {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}
	
	return duplicates
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func insertBST(root *TreeNode, val int, freq map[int]int) *TreeNode {
	if root == nil {
		freq[val]++
		return &TreeNode{val: val}
	}
	
	if val < root.val {
		root.left = insertBST(root.left, val, freq)
	} else if val > root.val {
		root.right = insertBST(root.right, val, freq)
	} else {
		freq[val]++
	}
	
	return root
}

// Divide and conquer with partitioning
func findDuplicatesPartition(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Choose pivot and partition
	pivot := nums[len(nums)/2]
	left, right := partitionByPivot(nums, pivot)
	
	// Recursively find duplicates in each partition
	leftDuplicates := findDuplicatesPartition(left)
	rightDuplicates := findDuplicatesPartition(right)
	
	// Check pivot itself
	pivotCount := 0
	for _, num := range nums {
		if num == pivot {
			pivotCount++
		}
	}
	
	// Merge results
	duplicateMap := make(map[int]bool)
	for _, dup := range leftDuplicates {
		duplicateMap[dup] = true
	}
	for _, dup := range rightDuplicates {
		duplicateMap[dup] = true
	}
	if pivotCount > 1 {
		duplicateMap[pivot] = true
	}
	
	var result []int
	for dup := range duplicateMap {
		result = append(result, dup)
	}
	
	return result
}

func partitionByPivot(nums []int, pivot int) ([]int, []int) {
	var left, right []int
	
	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		}
	}
	
	return left, right
}

// Divide and conquer with merge and find
func findDuplicatesMergeFind(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	
	// Sort and merge
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	mergeSort(sorted, 0, len(sorted)-1)
	
	// Find duplicates using divide and conquer
	return findDuplicatesInSorted(sorted, 0, len(sorted)-1)
}

func findDuplicatesInSorted(nums []int, left, right int) []int {
	if left >= right {
		return []int{}
	}
	
	if left == right {
		return []int{}
	}
	
	mid := left + (right-left)/2
	
	leftDuplicates := findDuplicatesInSorted(nums, left, mid)
	rightDuplicates := findDuplicatesInSorted(nums, mid+1, right)
	
	// Check for duplicates that cross the boundary
	crossDuplicates := findCrossDuplicates(nums, left, mid, right)
	
	// Merge and deduplicate
	duplicateMap := make(map[int]bool)
	
	for _, dup := range leftDuplicates {
		duplicateMap[dup] = true
	}
	for _, dup := range rightDuplicates {
		duplicateMap[dup] = true
	}
	for _, dup := range crossDuplicates {
		duplicateMap[dup] = true
	}
	
	var result []int
	for dup := range duplicateMap {
		result = append(result, dup)
	}
	
	return result
}

func findCrossDuplicates(nums []int, left, mid, right int) []int {
	var duplicates []int
	
	// Check if nums[mid] == nums[mid+1]
	if mid+1 <= right && nums[mid] == nums[mid+1] {
		duplicates = append(duplicates, nums[mid])
	}
	
	return duplicates
}

// Optimized O(1) space solution (for comparison)
func findDuplicatesOptimized(nums []int) []int {
	var duplicates []int
	
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		
		if nums[index] < 0 {
			duplicates = append(duplicates, abs(nums[i]))
		} else {
			nums[index] = -nums[index]
		}
	}
	
	return duplicates
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Divide and Conquer for Duplicate Detection
- **Sort and Scan**: Sort array using divide and conquer, then scan for duplicates
- **Frequency Counting**: Divide array to build frequency maps recursively
- **Partitioning**: Split by pivot values and find duplicates in partitions
- **Merge and Find**: Sort first, then use divide and conquer to find duplicates

## 2. PROBLEM CHARACTERISTICS
- **Duplicate Detection**: Find all elements that appear more than once
- **Array Processing**: Work with integer arrays of various sizes
- **Frequency Analysis**: Need to count occurrences of each element
- **Result Collection**: Gather unique duplicate values

## 3. SIMILAR PROBLEMS
- Find All Duplicates in an Array (LeetCode 442) - Same problem
- Contains Duplicate - Check if any duplicates exist
- Top K Frequent Elements - Find most frequent elements
- Find Missing Numbers - Complementary problem

## 4. KEY OBSERVATIONS
- **Sorting Natural**: After sorting, duplicates become adjacent
- **Frequency Maps**: Counting occurrences reveals duplicates
- **Divide Applicable**: Can split problem into smaller subarrays
- **Merge Strategy**: Combine results from subproblems

## 5. VARIATIONS & EXTENSIONS
- **Sort and Scan**: O(N log N) time, O(N) space (for sorting)
- **Hash Map**: O(N) time, O(N) space - optimal
- **Counting Sort**: O(N + k) time, O(k) space for limited range
- **In-Place Marking**: O(N) time, O(1) space - modifies input

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size? Value range? Can modify input? Space constraints?"
- Edge cases: empty array, single element, all duplicates, no duplicates
- Time complexity: O(N log N) for sorting, O(N) for hash map
- Space complexity: O(N) for hash map, O(1) for in-place
- Key insight: sorting makes duplicate detection trivial

## 7. COMMON MISTAKES
- Not handling multiple occurrences of same duplicate
- Wrong merge sort implementation boundaries
- Missing base cases in recursive divide and conquer
- Not deduplicating results properly
- Incorrect partitioning logic

## 8. OPTIMIZATION STRATEGIES
- **Sort and Scan**: O(N log N) time, O(N) space - demonstrates D&C
- **Hash Map**: O(N) time, O(N) space - optimal
- **Counting Sort**: O(N + k) time, O(k) space - for limited range
- **In-Place Marking**: O(N) time, O(1) space - modifies input

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing a library and finding duplicate books:**
- You have a pile of books with some duplicates
- You want to identify which books appear more than once
- You can sort the books by title first, then scan for duplicates
- Or you can create a catalog counting each book's occurrences
- Like a librarian organizing books and flagging duplicates for removal

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (may contain duplicates)
2. **Goal**: Find all values that appear more than once
3. **Constraints**: Return each duplicate only once
4. **Output**: Array of duplicate values (order doesn't matter)

#### Phase 2: Key Insight Recognition
- **"Sorting helps"** → After sorting, duplicates become adjacent
- **"Counting natural"** → Frequency counting reveals duplicates
- **"Divide applicable"** → Can split array and process subarrays
- **"Merge results"** → Combine duplicate findings from subproblems

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find all duplicate elements.
Brute force: compare every pair O(N²).

Divide and Conquer Approach:
1. Sort array using merge sort (divide and conquer)
2. Scan sorted array for adjacent equal elements
3. Collect each duplicate once

Alternative: Hash Map
1. Build frequency map in O(N) time
2. Extract keys with count > 1

Sort and Scan: O(N log N) time, O(N) space
Hash Map: O(N) time, O(N) space - better!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single element**: No duplicates possible
- **All duplicates**: Return unique values
- **No duplicates**: Return empty array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [4, 3, 2, 7, 8, 2, 3, 1]

Human thinking:
"Sort and Scan Process:
Step 1: Sort array using merge sort
Original: [4, 3, 2, 7, 8, 2, 3, 1]
Sorted:   [1, 2, 2, 3, 3, 4, 7, 8]

Step 2: Scan for duplicates
Index 0: [1] - no previous, continue
Index 1: [1, 2] - 2 != 1, continue
Index 2: [1, 2, 2] - 2 == 2, found duplicate!
Add 2 to result, skip additional 2s
Index 3: [1, 2, 2, 3] - 3 != 2, continue
Index 4: [1, 2, 2, 3, 3] - 3 == 3, found duplicate!
Add 3 to result, skip additional 3s
Index 5: [1, 2, 2, 3, 3, 4] - 4 != 3, continue
Index 6: [1, 2, 2, 3, 3, 4, 7] - 7 != 4, continue
Index 7: [1, 2, 2, 3, 3, 4, 7, 8] - 8 != 7, continue

Result: [2, 3] ✓"
```

#### Phase 6: Intuition Validation
- **Why sorting**: Makes duplicates adjacent and easy to detect
- **Why divide and conquer**: Merge sort naturally uses divide and conquer
- **Why hash map better**: O(N) time vs O(N log N) for sorting
- **Why single pass**: After sorting, one scan finds all duplicates

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use hash map?"** → Sorting approach demonstrates divide and conquer pattern
2. **"Should I check all pairs?"** → O(N²) vs O(N log N), too slow
3. **"What about multiple duplicates?"** → Need to skip additional occurrences
4. **"Can I modify the array?"** → In-place solutions possible but modify input
5. **"Why merge sort?"** → Stable, O(N log N), uses divide and conquer

### Real-World Analogy
**Like finding duplicate student IDs in a school system:**
- You have a list of student IDs with some duplicates
- You want to identify which IDs were assigned to multiple students
- You can sort the IDs first, then scan for duplicates
- Or you can create a hash table counting each ID's occurrences
- Like a registrar finding duplicate ID assignments for correction

### Human-Readable Pseudocode
```
function findDuplicates(nums):
    if len(nums) <= 1:
        return []
    
    # Sort using merge sort (divide and conquer)
    sorted = mergeSort(nums)
    
    # Scan for duplicates
    duplicates = []
    for i from 1 to len(sorted)-1:
        if sorted[i] == sorted[i-1]:
            duplicates.append(sorted[i])
            # Skip additional occurrences
            while i+1 < len(sorted) and sorted[i+1] == sorted[i]:
                i += 1
    
    return duplicates

function mergeSort(nums):
    if len(nums) <= 1:
        return nums
    
    mid = len(nums) // 2
    left = mergeSort(nums[:mid])
    right = mergeSort(nums[mid:])
    
    return merge(left, right)
```

### Execution Visualization

### Example: nums = [4, 3, 2, 7, 8, 2, 3, 1]
```
Merge Sort Process:
Level 0: [4, 3, 2, 7, 8, 2, 3, 1]
         Split at index 4
    Left: [4, 3, 2, 7]    Right: [8, 2, 3, 1]

Level 1: Left: [4, 3, 2, 7]
         Split at index 2
    LeftLeft: [4, 3]    LeftRight: [2, 7]

Level 2: LeftLeft: [4, 3]
         Split at index 1
    [4]    [3]
    Merge: [3, 4]

Level 2: LeftRight: [2, 7]
         Split at index 3
    [2]    [7]
    Merge: [2, 7]

Level 1: Left: [3, 4, 2, 7]
    Merge: [2, 3, 4, 7]

Continue similar process for right half...
Final sorted: [1, 2, 2, 3, 3, 4, 7, 8]

Scan Process:
[1] - no duplicate
[1, 2] - no duplicate  
[1, 2, 2] - duplicate found: 2
[1, 2, 2, 3] - no duplicate
[1, 2, 2, 3, 3] - duplicate found: 3
Continue...
Result: [2, 3] ✓
```

### Key Visualization Points:
- **Recursive Splitting**: Array split until single elements
- **Merge Process**: Combine sorted halves maintaining order
- **Duplicate Detection**: Adjacent equal elements in sorted array
- **Skip Logic**: Avoid adding same duplicate multiple times

### Merge Sort Tree Visualization:
```
        [4,3,2,7,8,2,3,1]
       /              \
    [4,3,2,7]        [8,2,3,1]
    /      \        /      \
 [4,3]    [2,7]  [8,2]    [3,1]
  /   \    /   \  /   \    /   \
[4]   [3][2]   [7][8]   [2][3]   [1]
```

### Time Complexity Breakdown:
- **Sort and Scan**: O(N log N) time, O(N) space - demonstrates D&C
- **Hash Map**: O(N) time, O(N) space - optimal
- **Counting Sort**: O(N + k) time, O(k) space - for limited range
- **In-Place Marking**: O(N) time, O(1) space - modifies input

### Alternative Approaches:

#### 1. Hash Map (O(N) time, O(N) space)
```go
func findDuplicatesHashMap(nums []int) []int {
    freq := make(map[int]int)
    
    for _, num := range nums {
        freq[num]++
    }
    
    var duplicates []int
    for num, count := range freq {
        if count > 1 {
            duplicates = append(duplicates, num)
        }
    }
    
    return duplicates
}
```
- **Pros**: Optimal O(N) time, simple implementation
- **Cons**: O(N) extra space

#### 2. In-Place Marking (O(N) time, O(1) space)
```go
func findDuplicatesInPlace(nums []int) []int {
    var duplicates []int
    
    for i := 0; i < len(nums); i++ {
        index := abs(nums[i]) - 1
        
        if nums[index] < 0 {
            duplicates = append(duplicates, abs(nums[i]))
        } else {
            nums[index] = -nums[index]
        }
    }
    
    return duplicates
}
```
- **Pros**: O(1) extra space, O(N) time
- **Cons**: Modifies input array, assumes values 1..N

#### 3. Counting Sort (O(N + k) time, O(k) space)
```go
func findDuplicatesCounting(nums []int) []int {
    minVal, maxVal := nums[0], nums[0]
    for _, num := range nums {
        if num < minVal { minVal = num }
        if num > maxVal { maxVal = num }
    }
    
    rangeSize := maxVal - minVal + 1
    count := make([]int, rangeSize)
    
    for _, num := range nums {
        count[num-minVal]++
    }
    
    var duplicates []int
    for i, freq := range count {
        if freq > 1 {
            duplicates = append(duplicates, i+minVal)
        }
    }
    
    return duplicates
}
```
- **Pros**: Linear time for limited range
- **Cons**: Space depends on value range

### Extensions for Interviews:
- **Multiple Duplicates**: Return all occurrences, not just unique values
- **Range Constraints**: Handle specific value ranges efficiently
- **Streaming Data**: Find duplicates in data stream
- **Memory Constraints**: Find duplicates with limited memory
- **Real-world Applications**: Data cleaning, fraud detection, inventory management
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Find All Duplicates in an Array ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, "Standard case"},
		{[]int{1, 2, 3, 4, 5}, "No duplicates"},
		{[]int{1, 1, 2, 2, 3, 3}, "All duplicates"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
		{[]int{2, 2, 2, 2}, "All same"},
		{[]int{1, 2, 3, 1, 4, 2}, "Some duplicates"},
		{[]int{5, 4, 3, 2, 1, 5, 4, 3}, "Reverse with duplicates"},
		{[]int{100, 200, 100, 300, 200, 400}, "Large numbers"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2}, "Long array"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Input: %v\n", tc.nums)
		
		// Make copies for different algorithms
		copy1 := make([]int, len(tc.nums))
		copy(copy1, tc.nums)
		copy2 := make([]int, len(tc.nums))
		copy(copy2, tc.nums)
		copy3 := make([]int, len(tc.nums))
		copy(copy3, tc.nums)
		copy4 := make([]int, len(tc.nums))
		copy(copy4, tc.nums)
		copy5 := make([]int, len(tc.nums))
		copy(copy5, tc.nums)
		
		result1 := findDuplicates(copy1)
		result2 := findDuplicatesCountingSort(copy2)
		result3 := findDuplicatesHashMap(copy3)
		result4 := findDuplicatesBST(copy4)
		result5 := findDuplicatesPartition(copy5)
		
		fmt.Printf("  Sort & Scan: %v\n", result1)
		fmt.Printf("  Counting Sort: %v\n", result2)
		fmt.Printf("  Hash Map: %v\n", result3)
		fmt.Printf("  BST: %v\n", result4)
		fmt.Printf("  Partition: %v\n\n", result5)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = (i % 1000) + 1
	}
	
	fmt.Printf("Large array test with %d elements\n", len(largeArray))
	
	// Test different approaches
	copy1 := make([]int, len(largeArray))
	copy(copy1, largeArray)
	result := findDuplicates(copy1)
	fmt.Printf("Sort & Scan found %d duplicates\n", len(result))
	
	copy2 := make([]int, len(largeArray))
	copy(copy2, largeArray)
	result = findDuplicatesHashMap(copy2)
	fmt.Printf("Hash Map found %d duplicates\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Array with all duplicates
	allDup := []int{5, 5, 5, 5, 5}
	fmt.Printf("All duplicates: %v\n", findDuplicates(allDup))
	
	// Array with negative numbers
	negatives := []int{-1, -2, -1, -3, -2}
	fmt.Printf("With negatives: %v\n", findDuplicates(negatives))
	
	// Array with zeros
	withZeros := []int{0, 1, 0, 2, 3, 0}
	fmt.Printf("With zeros: %v\n", findDuplicates(withZeros))
	
	// Very large numbers
	veryLarge := []int{1000000, 2000000, 1000000, 3000000}
	fmt.Printf("Very large numbers: %v\n", findDuplicates(veryLarge))
	
	// Test optimized solution for comparison
	fmt.Println("\n=== Optimized Solution Test ===")
	testArray := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Printf("Optimized O(1) space: %v\n", findDuplicatesOptimized(testArray))
	
	// Test merge find approach
	fmt.Println("\n=== Merge Find Test ===")
	mergeFindResult := findDuplicatesMergeFind(testArray)
	fmt.Printf("Merge Find: %v\n", mergeFindResult)
	
	// Test with many small duplicates
	fmt.Println("\n=== Many Small Duplicates Test ===")
	manySmall := make([]int, 1000)
	for i := range manySmall {
		manySmall[i] = (i % 10) + 1
	}
	
	result = findDuplicates(manySmall)
	fmt.Printf("Many small duplicates: %v (should be 1-10)\n", result)
}
