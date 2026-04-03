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
