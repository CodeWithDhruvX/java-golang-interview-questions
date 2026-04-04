package main

import (
	"fmt"
	"math"
)

// 912. Sort an Array - Divide and Conquer (Merge Sort)
// Time: O(N log N), Space: O(N)
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	
	// Create temporary array for merging
	temp := make([]int, len(nums))
	mergeSort(nums, temp, 0, len(nums)-1)
	return nums
}

func mergeSort(nums, temp []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		
		// Sort left half
		mergeSort(nums, temp, left, mid)
		
		// Sort right half
		mergeSort(nums, temp, mid+1, right)
		
		// Merge sorted halves
		merge(nums, temp, left, mid, right)
	}
}

func merge(nums, temp []int, left, mid, right int) {
	// Copy to temporary array
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}
	
	i := left     // Index for left subarray
	j := mid + 1  // Index for right subarray
	k := left     // Index for merged array
	
	// Merge while both subarrays have elements
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements from left subarray
	for i <= mid {
		nums[k] = temp[i]
		i++
		k++
	}
	
	// Copy remaining elements from right subarray
	for j <= right {
		nums[k] = temp[j]
		j++
		k++
	}
}

// In-place merge sort (bottom-up)
func sortArrayBottomUp(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	
	n := len(nums)
	temp := make([]int, n)
	
	// Start with subarrays of size 1 and double each time
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := min(left+size-1, n-1)
			right := min(left+2*size-1, n-1)
			
			if mid < right {
				merge(nums, temp, left, mid, right)
			}
		}
	}
	
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Quick Sort (another divide and conquer approach)
func sortArrayQuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low < high {
		// Partition the array
		pi := partition(nums, low, high)
		
		// Sort elements before and after partition
		quickSort(nums, low, pi-1)
		quickSort(nums, pi+1, high)
	}
}

func partition(nums []int, low, high int) int {
	// Choose the rightmost element as pivot
	pivot := nums[high]
	
	// Index of smaller element
	i := low - 1
	
	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	
	// Place pivot at the correct position
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

// Heap Sort (divide and conquer variant)
func sortArrayHeapSort(nums []int) []int {
	n := len(nums)
	
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	
	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		nums[0], nums[i] = nums[i], nums[0]
		
		// Call heapify on the reduced heap
		heapify(nums, i, 0)
	}
	
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left child
	right := 2*i + 2 // right child
	
	// If left child is larger than root
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	
	// If right child is larger than largest so far
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	
	// If largest is not root
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		
		// Recursively heapify the affected sub-tree
		heapify(nums, n, largest)
	}
}

// Merge Sort with counting inversions
func sortArrayCountInversions(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}
	
	temp := make([]int, len(nums))
	inversions := mergeSortCount(nums, temp, 0, len(nums)-1)
	return nums, inversions
}

func mergeSortCount(nums, temp []int, left, right int) int {
	if left >= right {
		return 0
	}
	
	mid := left + (right-left)/2
	
	inversions := mergeSortCount(nums, temp, left, mid)
	inversions += mergeSortCount(nums, temp, mid+1, right)
	inversions += mergeCount(nums, temp, left, mid, right)
	
	return inversions
}

func mergeCount(nums, temp []int, left, mid, right int) int {
	// Copy to temporary array
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}
	
	i := left     // Index for left subarray
	j := mid + 1  // Index for right subarray
	k := left     // Index for merged array
	inversions := 0
	
	// Merge while both subarrays have elements
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
			// All remaining elements in left subarray are greater than temp[j]
			inversions += (mid - i + 1)
		}
		k++
	}
	
	// Copy remaining elements
	for i <= mid {
		nums[k] = temp[i]
		i++
		k++
	}
	
	for j <= right {
		nums[k] = temp[j]
		j++
		k++
	}
	
	return inversions
}

// External merge sort simulation (for very large arrays)
func sortArrayExternal(nums []int) []int {
	if len(nums) <= 1000 {
		return sortArray(nums)
	}
	
	// Simulate external sort by dividing into chunks
	chunkSize := 1000
	var chunks [][]int
	
	// Divide into chunks and sort each
	for i := 0; i < len(nums); i += chunkSize {
		end := min(i+chunkSize, len(nums))
		chunk := make([]int, end-i)
		copy(chunk, nums[i:end])
		sortArray(chunk)
		chunks = append(chunks, chunk)
	}
	
	// Merge chunks
	result := make([]int, 0, len(nums))
	
	// Simple k-way merge using priority queue simulation
	for len(chunks) > 0 {
		minIdx := 0
		for i := 1; i < len(chunks); i++ {
			if len(chunks[i]) > 0 && chunks[i][0] < chunks[minIdx][0] {
				minIdx = i
			}
		}
		
		if len(chunks[minIdx]) > 0 {
			result = append(result, chunks[minIdx][0])
			chunks[minIdx] = chunks[minIdx][1:]
			
			if len(chunks[minIdx]) == 0 {
				chunks = append(chunks[:minIdx], chunks[minIdx+1:]...)
			}
		}
	}
	
	return result
}

// Merge Sort with stability check
func sortArrayStable(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	
	// Create array of value-index pairs to track stability
	type Element struct {
		value int
		index int
	}
	
	elements := make([]Element, len(nums))
	for i, val := range nums {
		elements[i] = Element{val, i}
	}
	
	temp := make([]Element, len(elements))
	mergeSortStable(elements, temp, 0, len(elements)-1)
	
	// Extract values
	result := make([]int, len(nums))
	for i, elem := range elements {
		result[i] = elem.value
	}
	
	return result
}

func mergeSortStable(elements, temp []Element, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortStable(elements, temp, left, mid)
		mergeSortStable(elements, temp, mid+1, right)
		mergeStable(elements, temp, left, mid, right)
	}
}

func mergeStable(elements, temp []Element, left, mid, right int) {
	for i := left; i <= right; i++ {
		temp[i] = elements[i]
	}
	
	i := left
	j := mid + 1
	k := left
	
	for i <= mid && j <= right {
		if temp[i].value <= temp[j].value {
			elements[k] = temp[i]
			i++
		} else {
			elements[k] = temp[j]
			j++
		}
		k++
	}
	
	for i <= mid {
		elements[k] = temp[i]
		i++
		k++
	}
	
	for j <= right {
		elements[k] = temp[j]
		j++
		k++
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Divide and Conquer for Array Sorting
- **Merge Sort**: Split array, sort halves, merge sorted results
- **Quick Sort**: Partition around pivot, sort subarrays recursively
- **Heap Sort**: Build heap, extract elements in sorted order
- **External Sort**: Handle large arrays with chunk-based processing

## 2. PROBLEM CHARACTERISTICS
- **Array Sorting**: Rearrange elements in non-decreasing order
- **Stability Requirement**: Maintain relative order of equal elements
- **Memory Constraints**: Different space complexity requirements
- **Performance Needs**: Time complexity optimization for different inputs

## 3. SIMILAR PROBLEMS
- Sort an Array (LeetCode 912) - Same problem
- Sort List - Linked list sorting
- Sort Colors - Limited range sorting
- Top K Elements - Partial sorting

## 4. KEY OBSERVATIONS
- **Divide Natural**: Array can be split at midpoint
- **Merge Strategy**: Combine sorted subarrays efficiently
- **Partition Strategy**: Divide elements around pivot value
- **Stability**: Merge sort maintains stability naturally

## 5. VARIATIONS & EXTENSIONS
- **Merge Sort**: O(N log N) time, O(N) space - stable
- **Quick Sort**: O(N log N) average, O(N²) worst - in-place
- **Heap Sort**: O(N log N) time, O(1) space - not stable
- **External Sort**: Handle arrays larger than memory

## 6. INTERVIEW INSIGHTS
- Always clarify: "Stability required? Space constraints? Input size?"
- Edge cases: empty array, single element, all duplicates, already sorted
- Time complexity: O(N log N) for merge/heap, O(N log N) avg for quick
- Space complexity: O(N) for merge, O(log N) for quick, O(1) for heap
- Key insight: merge sort stable, quick sort in-place, heap sort space-efficient

## 7. COMMON MISTAKES
- Wrong merge sort boundaries and indexing
- Incorrect quick sort partition implementation
- Missing base cases in recursive functions
- Not handling empty arrays properly
- Wrong heap construction in heap sort

## 8. OPTIMIZATION STRATEGIES
- **Merge Sort**: O(N log N) time, O(N) space - stable, predictable
- **Quick Sort**: O(N log N) time, O(log N) space - fast in practice
- **Heap Sort**: O(N log N) time, O(1) space - space efficient
- **Bottom-up Merge**: O(N log N) time, O(N) space - iterative

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing a deck of cards:**
- You have a shuffled deck of cards (unsorted array)
- You want to arrange them in order (sorted array)
- You can split the deck into smaller piles, sort each pile, then merge
- Or you can pick a card as reference and separate others around it
- Like a card player organizing cards by splitting and combining

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers in arbitrary order
2. **Goal**: Rearrange elements in non-decreasing order
3. **Constraints**: May need stability, space efficiency
4. **Output**: Same array elements in sorted order

#### Phase 2: Key Insight Recognition
- **"Divide natural"** → Can split array at midpoint
- **"Merge efficient"** → Can combine two sorted arrays efficiently
- **"Partition effective"** → Can separate elements around pivot
- **"Recursive structure"** → Same problem applied to subarrays

#### Phase 3: Strategy Development
```
Human thought process:
"I need to sort an array.
Brute force: bubble sort O(N²), too slow.

Divide and Conquer Approach:
1. Split array at midpoint recursively
2. Sort each half (same problem, smaller size)
3. Merge sorted halves efficiently

Merge Sort: O(N log N) time, O(N) space, stable
Quick Sort: O(N log N) avg, O(log N) space, in-place
Heap Sort: O(N log N) time, O(1) space, not stable

Each has different trade-offs!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return as-is
- **Single element**: Already sorted
- **All duplicates**: Maintain stability if required
- **Already sorted**: Algorithm handles efficiently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [5, 2, 3, 1]

Human thinking:
"Merge Sort Process:
Step 1: Split at middle (index 1)
Left: [5, 2], Right: [3, 1]

Step 2: Recursively sort left [5, 2]
Split at index 0
[5] and [2] - base cases, already sorted
Merge [5] and [2]: [2, 5]

Step 3: Recursively sort right [3, 1]
Split at index 2
[3] and [1] - base cases, already sorted
Merge [3] and [1]: [1, 3]

Step 4: Merge [2, 5] and [1, 3]
Compare 2 vs 1: take 1
Compare 2 vs 3: take 2
Compare 5 vs 3: take 3
Take remaining 5
Result: [1, 2, 3, 5] ✓"
```

#### Phase 6: Intuition Validation
- **Why divide**: Natural way to break down problem
- **Why merge**: Efficiently combine sorted subarrays
- **Why O(N log N)**: Each level does O(N) work, O(log N) levels
- **Why stability**: Merge preserves relative order of equal elements

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use built-in sort?"** → Need to understand algorithm implementation
2. **"Should I use bubble sort?"** → O(N²) vs O(N log N), too slow
3. **"What about insertion sort?"** → O(N²) worst, good for small arrays
4. **"Can I sort in-place?"** → Quick sort and heap sort are in-place
5. **"Why multiple algorithms?"** → Different trade-offs for different scenarios

### Real-World Analogy
**Like organizing books in a library:**
- You have books in random order on shelves (unsorted array)
- You want to arrange them by call number (sorted array)
- You can split books into sections, sort each section, then merge
- Or you can pick a reference book and organize others around it
- Like a librarian organizing books by section and then merging sections

### Human-Readable Pseudocode
```
function mergeSort(nums):
    if len(nums) <= 1:
        return nums
    
    mid = len(nums) // 2
    left = mergeSort(nums[:mid])
    right = mergeSort(nums[mid:])
    
    return merge(left, right)

function merge(left, right):
    result = []
    i = j = 0
    
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    
    # Add remaining elements
    result.extend(left[i:])
    result.extend(right[j:])
    
    return result
```

### Execution Visualization

### Example: nums = [5, 2, 3, 1]
```
Merge Sort Process:
Level 0: [5, 2, 3, 1]
         Split at index 1
    Left: [5, 2]    Right: [3, 1]

Level 1: Left: [5, 2]
         Split at index 0
    [5]    [2]
    Merge: [2, 5]

Level 1: Right: [3, 1]
         Split at index 2
    [3]    [1]
    Merge: [1, 3]

Level 0: Merge [2, 5] and [1, 3]
[2, 5] + [1, 3] → [1, 2, 3, 5]

Final result: [1, 2, 3, 5] ✓
```

### Key Visualization Points:
- **Recursive Splitting**: Array split until single elements
- **Merge Process**: Combine sorted halves maintaining order
- **Stability**: Equal elements maintain relative order
- **Efficiency**: Linear merge at each level

### Merge Sort Tree Visualization:
```
        [5,2,3,1]
       /          \
    [5,2]        [3,1]
    /   \        /   \
[5]   [2]    [3]   [1]
  \   /        \   /
   [2,5]        [1,3]
      \          /
       [1,2,3,5]
```

### Time Complexity Breakdown:
- **Merge Sort**: O(N log N) time, O(N) space - stable, predictable
- **Quick Sort**: O(N log N) average, O(N²) worst, O(log N) space - fast
- **Heap Sort**: O(N log N) time, O(1) space - space efficient
- **Bottom-up Merge**: O(N log N) time, O(N) space - iterative

### Alternative Approaches:

#### 1. Quick Sort (O(N log N) average, O(log N) space)
```go
func quickSort(nums []int, low, high int) {
    if low < high {
        pi := partition(nums, low, high)
        quickSort(nums, low, pi-1)
        quickSort(nums, pi+1, high)
    }
}
```
- **Pros**: Fast in practice, in-place, good cache performance
- **Cons**: O(N²) worst case, not stable

#### 2. Heap Sort (O(N log N) time, O(1) space)
```go
func heapSort(nums []int) {
    n := len(nums)
    
    // Build max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(nums, n, i)
    }
    
    // Extract elements
    for i := n - 1; i > 0; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        heapify(nums, i, 0)
    }
}
```
- **Pros**: O(1) space, O(N log N) guaranteed
- **Cons**: Not stable, poor cache performance

#### 3. Built-in Sort (O(N log N) time, varies by language)
```go
func sortArray(nums []int) []int {
    sort.Ints(nums)
    return nums
}
```
- **Pros**: Simple, optimized implementation
- **Cons**: Black box, learning opportunity lost

### Extensions for Interviews:
- **Inversion Counting**: Count pairs out of order during merge
- **External Sorting**: Handle arrays larger than memory
- **Stability Requirements**: Maintain relative order of equal elements
- **Parallel Sorting**: Use multiple threads for sorting
- **Real-world Applications**: Database sorting, file systems, data processing
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Sort Array - Divide and Conquer ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{5, 2, 3, 1}, "Standard case"},
		{[]int{5, 1, 1, 2, 3, 5, 4}, "With duplicates"},
		{[]int{1, 2, 3, 4, 5}, "Already sorted"},
		{[]int{5, 4, 3, 2, 1}, "Reverse sorted"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
		{[]int{2, 2, 2, 2}, "All same"},
		{[]int{-1, -3, -2, -5, -4}, "All negative"},
		{[]int{0, 0, 0, 0}, "All zeros"},
		{[]int{100, -100, 50, -50, 0}, "Mixed positive and negative"},
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
		
		result1 := sortArray(copy1)
		result2 := sortArrayBottomUp(copy2)
		result3 := sortArrayQuickSort(copy3)
		result4 := sortArrayHeapSort(copy4)
		result5 := sortArrayStable(copy5)
		
		fmt.Printf("  Merge Sort: %v\n", result1)
		fmt.Printf("  Bottom-up: %v\n", result2)
		fmt.Printf("  Quick Sort: %v\n", result3)
		fmt.Printf("  Heap Sort: %v\n", result4)
		fmt.Printf("  Stable: %v\n\n", result5)
	}
	
	// Test inversion counting
	fmt.Println("=== Inversion Counting Test ===")
	testArray := []int{2, 4, 1, 3, 5}
	sorted, inversions := sortArrayCountInversions(testArray)
	fmt.Printf("Original: %v\n", testArray)
	fmt.Printf("Sorted: %v\n", sorted)
	fmt.Printf("Inversions: %d\n", inversions)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = 10000 - i // Reverse order
	}
	
	fmt.Printf("Large array test with %d elements\n", len(largeArray))
	
	// Test merge sort
	copy1 := make([]int, len(largeArray))
	copy(copy1, largeArray)
	result := sortArray(copy1)
	fmt.Printf("Merge sort completed\n")
	
	// Test quick sort
	copy2 := make([]int, len(largeArray))
	copy(copy2, largeArray)
	result = sortArrayQuickSort(copy2)
	fmt.Printf("Quick sort completed\n")
	
	// Test heap sort
	copy3 := make([]int, len(largeArray))
	copy(copy3, largeArray)
	result = sortArrayHeapSort(copy3)
	fmt.Printf("Heap sort completed\n")
	
	// Test external sort simulation
	copy4 := make([]int, len(largeArray))
	copy(copy4, largeArray)
	result = sortArrayExternal(copy4)
	fmt.Printf("External sort simulation completed\n")
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Already sorted large array
	sortedLarge := make([]int, 5000)
	for i := range sortedLarge {
		sortedLarge[i] = i
	}
	
	copy1 = make([]int, len(sortedLarge))
	copy(copy1, sortedLarge)
	result = sortArray(copy1)
	fmt.Printf("Already sorted large array: first few elements %v\n", result[:5])
	
	// Array with many duplicates
	duplicates := make([]int, 5000)
	for i := range duplicates {
		duplicates[i] = i % 100
	}
	
	copy1 = make([]int, len(duplicates))
	copy(copy1, duplicates)
	result = sortArray(copy1)
	fmt.Printf("Many duplicates: first few elements %v\n", result[:10])
	
	// Very large values
	largeVals := make([]int, 1000)
	for i := range largeVals {
		largeVals[i] = 1000000 + (i % 1000)
	}
	
	copy1 = make([]int, len(largeVals))
	copy(copy1, largeVals)
	result = sortArray(copy1)
	fmt.Printf("Large values: first few elements %v\n", result[:5])
}
