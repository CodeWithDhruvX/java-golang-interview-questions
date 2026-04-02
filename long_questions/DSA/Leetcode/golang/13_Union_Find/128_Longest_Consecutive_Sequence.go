package main

import "fmt"

// 128. Longest Consecutive Sequence
// Time: O(N), Space: O(N)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	
	maxLength := 0
	
	for num := range numSet {
		// Only start counting from the beginning of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1
			
			// Count the length of the consecutive sequence
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}
			
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	
	return maxLength
}

// Union-Find approach
func longestConsecutiveUnionFind(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Initialize Union-Find
	parent := make(map[int]int)
	rank := make(map[int]int)
	
	for _, num := range nums {
		parent[num] = num
		rank[num] = 0
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	var union func(int, int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return
		}
		
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
	}
	
	// Union consecutive numbers
	for _, num := range nums {
		if _, exists := parent[num-1]; exists {
			union(num, num-1)
		}
		if _, exists := parent[num+1]; exists {
			union(num, num+1)
		}
	}
	
	// Count the size of each component
	componentSize := make(map[int]int)
	for _, num := range nums {
		root := find(num)
		componentSize[root]++
	}
	
	maxLength := 0
	for _, size := range componentSize {
		if size > maxLength {
			maxLength = size
		}
	}
	
	return maxLength
}

// Sorting approach (O(N log N))
func longestConsecutiveSorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Remove duplicates and sort
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}
	
	sorted := make([]int, 0, len(unique))
	for num := range unique {
		sorted = append(sorted, num)
	}
	
	// Simple bubble sort for demonstration (in practice, use sort.Ints)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	maxLength := 1
	currentLength := 1
	
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1]+1 {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}
	
	if currentLength > maxLength {
		maxLength = currentLength
	}
	
	return maxLength
}

func main() {
	// Test cases
	testCases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		{},
		{1},
		{1, 2, 0, 1},
		{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 3, 5, 7, 9},
		{-1, -2, -3, 0, 1, 2, 3},
	}
	
	for i, nums := range testCases {
		result1 := longestConsecutive(nums)
		result2 := longestConsecutiveUnionFind(nums)
		result3 := longestConsecutiveSorting(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  HashSet: %d, Union-Find: %d, Sorting: %d\n\n", result1, result2, result3)
	}
}
