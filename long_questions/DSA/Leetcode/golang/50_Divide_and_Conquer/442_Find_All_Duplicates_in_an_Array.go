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
