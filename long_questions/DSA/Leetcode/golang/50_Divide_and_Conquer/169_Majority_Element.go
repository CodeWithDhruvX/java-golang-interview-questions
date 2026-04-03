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
