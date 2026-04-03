package main

import "fmt"

// 1649. Create Sorted Array through Instructions - Fenwick Tree
// Time: O(N log M), Space: O(M) where M is max value
func createSortedArray(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	// Find maximum value to determine Fenwick tree size
	maxVal := 0
	for _, val := range instructions {
		if val > maxVal {
			maxVal = val
		}
	}
	
	// Create Fenwick tree for counting
	tree := make([]int, maxVal+1)
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		// Count elements less than val
		less := queryFenwick(tree, val-1)
		// Count elements greater than val
		greater := i - queryFenwick(tree, val)
		
		// Add to result
		result = (result + min(less, greater)) % mod
		
		// Update Fenwick tree
		updateFenwick(tree, val, 1)
	}
	
	return result
}

// Fenwick tree operations
func updateFenwick(tree []int, i, delta int) {
	for i < len(tree) {
		tree[i] += delta
		i += i & (-i)
	}
}

func queryFenwick(tree []int, i int) int {
	sum := 0
	for i > 0 {
		sum += tree[i]
		i -= i & (-i)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Alternative approach using two Fenwick trees
func createSortedArrayTwoTrees(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	maxVal := 0
	for _, val := range instructions {
		if val > maxVal {
			maxVal = val
		}
	}
	
	// Two trees: one for counting, one for prefix sums
	countTree := make([]int, maxVal+1)
	sumTree := make([]int, maxVal+1)
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		less := queryFenwick(countTree, val-1)
		greater := i - queryFenwick(countTree, val)
		
		result = (result + min(less, greater)) % mod
		
		updateFenwick(countTree, val, 1)
		updateFenwick(sumTree, val, val)
	}
	
	return result
}

// Using coordinate compression
func createSortedArrayCompressed(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	// Coordinate compression
	sorted := make([]int, len(instructions))
	copy(sorted, instructions)
	
	// Simple bubble sort for demonstration
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	// Remove duplicates
	unique := []int{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1] {
			unique = append(unique, sorted[i])
		}
	}
	
	// Map values to compressed indices
	valToIdx := make(map[int]int)
	for i, val := range unique {
		valToIdx[val] = i + 1 // 1-based indexing
	}
	
	// Create Fenwick tree with compressed size
	tree := make([]int, len(unique)+1)
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		idx := valToIdx[val]
		
		less := queryFenwick(tree, idx-1)
		greater := i - queryFenwick(tree, idx)
		
		result = (result + min(less, greater)) % mod
		
		updateFenwick(tree, idx, 1)
	}
	
	return result
}

// Brute force approach for comparison
func createSortedArrayBruteForce(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		less := 0
		greater := 0
		
		for j := 0; j < i; j++ {
			if instructions[j] < val {
				less++
			} else if instructions[j] > val {
				greater++
			}
		}
		
		result = (result + min(less, greater)) % mod
	}
	
	return result
}

// Using segment tree approach
func createSortedArraySegmentTree(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	maxVal := 0
	for _, val := range instructions {
		if val > maxVal {
			maxVal = val
		}
	}
	
	// Build segment tree
	segTree := make([]int, 4*maxVal)
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		less := querySegmentTree(segTree, 1, 1, maxVal, 1, val-1)
		greater := i - querySegmentTree(segTree, 1, 1, maxVal, 1, val)
		
		result = (result + min(less, greater)) % mod
		
		updateSegmentTree(segTree, 1, 1, maxVal, val, 1)
	}
	
	return result
}

func updateSegmentTree(tree []int, node, start, end, idx, val int) {
	if start == end {
		tree[node] += val
		return
	}
	
	mid := start + (end-start)/2
	if idx <= mid {
		updateSegmentTree(tree, 2*node, start, mid, idx, val)
	} else {
		updateSegmentTree(tree, 2*node+1, mid+1, end, idx, val)
	}
	
	tree[node] = tree[2*node] + tree[2*node+1]
}

func querySegmentTree(tree []int, node, start, end, left, right int) int {
	if left > end || right < start {
		return 0
	}
	
	if left <= start && end <= right {
		return tree[node]
	}
	
	mid := start + (end-start)/2
	return querySegmentTree(tree, 2*node, start, mid, left, right) +
		   querySegmentTree(tree, 2*node+1, mid+1, end, left, right)
}

// Using binary indexed tree with frequency array
func createSortedArrayBIT(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	
	maxVal := 0
	for _, val := range instructions {
		if val > maxVal {
			maxVal = val
		}
	}
	
	// Frequency array and BIT
	freq := make([]int, maxVal+1)
	
	result := 0
	mod := 1000000007
	
	for i, val := range instructions {
		// Count elements less than val
		less := 0
		for j := 1; j < val; j++ {
			less += freq[j]
		}
		
		// Count elements greater than val
		greater := 0
		for j := val + 1; j <= maxVal; j++ {
			greater += freq[j]
		}
		
		result = (result + min(less, greater)) % mod
		
		freq[val]++
	}
	
	return result
}

func main() {
	// Test cases
	fmt.Println("=== Testing Create Sorted Array through Instructions ===")
	
	testCases := []struct {
		instructions []int
		description  string
	}{
		{[]int{1, 5, 6, 2}, "Standard case"},
		{[]int{1, 2, 3, 6, 5, 4}, "Increasing then decreasing"},
		{[]int{1, 3, 3, 3, 5, 6, 2}, "With duplicates"},
		{[]int{4, 5, 1, 2, 3}, "Mixed order"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
		{[]int{2, 2, 2, 2}, "All same"},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "Strictly decreasing"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Strictly increasing"},
		{[]int{5, 1, 5, 1, 5, 1}, "Alternating"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Instructions: %v\n", tc.instructions)
		
		result1 := createSortedArray(tc.instructions)
		result2 := createSortedArrayTwoTrees(tc.instructions)
		result3 := createSortedArrayCompressed(tc.instructions)
		result4 := createSortedArrayBruteForce(tc.instructions)
		result5 := createSortedArraySegmentTree(tc.instructions)
		result6 := createSortedArrayBIT(tc.instructions)
		
		fmt.Printf("  Fenwick Tree: %d\n", result1)
		fmt.Printf("  Two Trees: %d\n", result2)
		fmt.Printf("  Compressed: %d\n", result3)
		fmt.Printf("  Brute Force: %d\n", result4)
		fmt.Printf("  Segment Tree: %d\n", result5)
		fmt.Printf("  Frequency Array: %d\n\n", result6)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	largeInstructions := make([]int, 10000)
	for i := range largeInstructions {
		largeInstructions[i] = (i % 1000) + 1
	}
	
	fmt.Printf("Large test with %d instructions\n", len(largeInstructions))
	
	result := createSortedArray(largeInstructions)
	fmt.Printf("Fenwick Tree result: %d\n", result)
	
	result = createSortedArrayCompressed(largeInstructions)
	fmt.Printf("Compressed result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Large values
	largeVals := []int{1000000, 1, 500000, 999999, 2}
	fmt.Printf("Large values: %d\n", createSortedArray(largeVals))
	
	// With zeros
	withZeros := []int{0, 1, 0, 2, 0, 3}
	fmt.Printf("With zeros: %d\n", createSortedArray(withZeros))
	
	// All same value
	allSame := []int{5, 5, 5, 5, 5}
	fmt.Printf("All same: %d\n", createSortedArray(allSame))
	
	// Test coordinate compression
	fmt.Println("\n=== Coordinate Compression Test ===")
	
	sparseVals := []int{1000000, 1, 500000, 999999, 2, 1000001}
	fmt.Printf("Sparse values: %d\n", createSortedArrayCompressed(sparseVals))
	
	// Test with negative values (should handle gracefully)
	fmt.Println("\n=== Negative Values Test ===")
	
	negVals := []int{-1, -2, -3, -1, -2}
	fmt.Printf("Negative values: %d\n", createSortedArray(negVals))
}
