package main

import (
	"fmt"
	"math"
	"sort"
)

// 307. Range Sum Query - Mutable - Mos Algorithm
// Time: O((N+Q) * sqrt(N)), Space: O(N + Q)
type RangeSumQuery struct {
	nums     []int
	blockSize int
}

func Constructor(nums []int) RangeSumQuery {
	n := len(nums)
	blockSize := int(math.Sqrt(float64(n))) + 1
	
	return RangeSumQuery{
		nums:     nums,
		blockSize: blockSize,
	}
}

func (r *RangeSumQuery) Update(index int, val int) {
	r.nums[index] = val
}

func (r *RangeSumQuery) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += r.nums[i]
	}
	return sum
}

// Mos Algorithm for multiple range queries
type MosQuery struct {
	left  int
	right int
	index int
}

func (r *RangeSumQuery) SumRangeMos(queries []MosQuery) []int {
	if len(queries) == 0 {
		return []int{}
	}
	
	// Sort queries using Mo's ordering
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / r.blockSize
		blockJ := queries[j].left / r.blockSize
		
		if blockI != blockJ {
			return blockI < blockJ
		}
		
		// If same block, sort by right
		if blockI%2 == 0 {
			return queries[i].right < queries[j].right
		}
		return queries[i].right > queries[j].right
	})
	
	// Process queries
	results := make([]int, len(queries))
	currentLeft := 0
	currentRight := -1
	currentSum := 0
	
	for _, query := range queries {
		// Expand to right
		for currentRight < query.right {
			currentRight++
			currentSum += r.nums[currentRight]
		}
		
		// Contract from right
		for currentRight > query.right {
			currentSum -= r.nums[currentRight]
			currentRight--
		}
		
		// Expand to left
		for currentLeft < query.left {
			currentSum -= r.nums[currentLeft]
			currentLeft++
		}
		
		// Contract from left
		for currentLeft > query.left {
			currentLeft--
			currentSum += r.nums[currentLeft]
		}
		
		results[query.index] = currentSum
	}
	
	return results
}

// Mos Algorithm with frequency counting
type MosQueryFreq struct {
	left  int
	right int
	index int
	k     int // for frequency queries
}

func (r *RangeSumQuery) FrequencyQueries(queries []MosQueryFreq) []int {
	if len(queries) == 0 {
		return []int{}
	}
	
	// Sort queries
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / r.blockSize
		blockJ := queries[j].left / r.blockSize
		
		if blockI != blockJ {
			return blockI < blockJ
		}
		
		if blockI%2 == 0 {
			return queries[i].right < queries[j].right
		}
		return queries[i].right > queries[j].right
	})
	
	// Process queries with frequency counting
	results := make([]int, len(queries))
	freq := make(map[int]int)
	currentLeft := 0
	currentRight := -1
	
	for _, query := range queries {
		// Expand to right
		for currentRight < query.right {
			currentRight++
			freq[r.nums[currentRight]]++
		}
		
		// Contract from right
		for currentRight > query.right {
			freq[r.nums[currentRight]]--
			if freq[r.nums[currentRight]] == 0 {
				delete(freq, r.nums[currentRight])
			}
			currentRight--
		}
		
		// Expand to left
		for currentLeft < query.left {
			freq[r.nums[currentLeft]]--
			if freq[r.nums[currentLeft]] == 0 {
				delete(freq, r.nums[currentLeft])
			}
			currentLeft++
		}
		
		// Contract from left
		for currentLeft > query.left {
			currentLeft--
			freq[r.nums[currentLeft]]++
		}
		
		// Answer frequency query
		results[query.index] = freq[query.k]
	}
	
	return results
}

// Mos Algorithm with distinct count
type MosQueryDistinct struct {
	left  int
	right int
	index int
}

func (r *RangeSumQuery) DistinctCountQueries(queries []MosQueryDistinct) []int {
	if len(queries) == 0 {
		return []int{}
	}
	
	// Sort queries
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / r.blockSize
		blockJ := queries[j].left / r.blockSize
		
		if blockI != blockJ {
			return blockI < blockJ
		}
		
		if blockI%2 == 0 {
			return queries[i].right < queries[j].right
		}
		return queries[i].right > queries[j].right
	})
	
	// Process queries
	results := make([]int, len(queries))
	freq := make(map[int]int)
	currentLeft := 0
	currentRight := -1
	distinctCount := 0
	
	for _, query := range queries {
		// Expand to right
		for currentRight < query.right {
			currentRight++
			if freq[r.nums[currentRight]] == 0 {
				distinctCount++
			}
			freq[r.nums[currentRight]]++
		}
		
		// Contract from right
		for currentRight > query.right {
			freq[r.nums[currentRight]]--
			if freq[r.nums[currentRight]] == 0 {
				distinctCount--
				delete(freq, r.nums[currentRight])
			}
			currentRight--
		}
		
		// Expand to left
		for currentLeft < query.left {
			freq[r.nums[currentLeft]]--
			if freq[r.nums[currentLeft]] == 0 {
				distinctCount--
				delete(freq, r.nums[currentLeft])
			}
			currentLeft++
		}
		
		// Contract from left
		for currentLeft > query.left {
			currentLeft--
			if freq[r.nums[currentLeft]] == 0 {
				distinctCount++
			}
			freq[r.nums[currentLeft]]++
		}
		
		results[query.index] = distinctCount
	}
	
	return results
}

// Mos Algorithm with mode queries
type MosQueryMode struct {
	left  int
	right int
	index int
}

func (r *RangeSumQuery) ModeQueries(queries []MosQueryMode) []int {
	if len(queries) == 0 {
		return []int{}
	}
	
	// Sort queries
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / r.blockSize
		blockJ := queries[j].left / r.blockSize
		
		if blockI != blockJ {
			return blockI < blockJ
		}
		
		if blockI%2 == 0 {
			return queries[i].right < queries[j].right
		}
		return queries[i].right > queries[j].right
	})
	
	// Process queries
	results := make([]int, len(queries))
	freq := make(map[int]int)
	currentLeft := 0
	currentRight := -1
	mode := 0
	modeCount := 0
	
	for _, query := range queries {
		// Expand to right
		for currentRight < query.right {
			currentRight++
			freq[r.nums[currentRight]]++
			if freq[r.nums[currentRight]] > modeCount {
				mode = r.nums[currentRight]
				modeCount = freq[r.nums[currentRight]]
			}
		}
		
		// Contract from right
		for currentRight > query.right {
			freq[r.nums[currentRight]]--
			if freq[r.nums[currentRight]] == 0 {
				delete(freq, r.nums[currentRight])
			}
			if freq[r.nums[currentRight]] == modeCount-1 && mode == r.nums[currentRight] {
				// Need to recalculate mode
				modeCount = 0
				for num, count := range freq {
					if count > modeCount {
						mode = num
						modeCount = count
					}
				}
			}
			currentRight--
		}
		
		// Expand to left
		for currentLeft < query.left {
			freq[r.nums[currentLeft]]--
			if freq[r.nums[currentLeft]] == 0 {
				delete(freq, r.nums[currentLeft])
			}
			if freq[r.nums[currentLeft]] == modeCount-1 && mode == r.nums[currentLeft] {
				// Need to recalculate mode
				modeCount = 0
				for num, count := range freq {
					if count > modeCount {
						mode = num
						modeCount = count
					}
				}
			}
			currentLeft++
		}
		
		// Contract from left
		for currentLeft > query.left {
			currentLeft--
			freq[r.nums[currentLeft]]++
			if freq[r.nums[currentLeft]] > modeCount {
				mode = r.nums[currentLeft]
				modeCount = freq[r.nums[currentLeft]]
			}
		}
		
		results[query.index] = mode
	}
	
	return results
}

// Mos Algorithm with range minimum query
type MosQueryMin struct {
	left  int
	right int
	index int
}

func (r *RangeSumQuery) MinQueries(queries []MosQueryMin) []int {
	if len(queries) == 0 {
		return []int{}
	}
	
	// Sort queries
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / r.blockSize
		blockJ := queries[j].left / r.blockSize
		
		if blockI != blockJ {
			return blockI < blockJ
		}
		
		if blockI%2 == 0 {
			return queries[i].right < queries[j].right
		}
		return queries[i].right > queries[j].right
	})
	
	// Process queries
	results := make([]int, len(queries))
	freq := make(map[int]int)
	currentLeft := 0
	currentRight := -1
	minVal := math.MaxInt32
	
	for _, query := range queries {
		// Expand to right
		for currentRight < query.right {
			currentRight++
			freq[r.nums[currentRight]]++
			if r.nums[currentRight] < minVal {
				minVal = r.nums[currentRight]
			}
		}
		
		// Contract from right
		for currentRight > query.right {
			freq[r.nums[currentRight]]--
			if freq[r.nums[currentRight]] == 0 {
				delete(freq, r.nums[currentRight])
				if r.nums[currentRight] == minVal {
					// Need to recalculate min
					minVal = math.MaxInt32
					for num := range freq {
						if num < minVal {
							minVal = num
						}
					}
				}
			}
			currentRight--
		}
		
		// Expand to left
		for currentLeft < query.left {
			freq[r.nums[currentLeft]]--
			if freq[r.nums[currentLeft]] == 0 {
				delete(freq, r.nums[currentLeft])
				if r.nums[currentLeft] == minVal {
					// Need to recalculate min
					minVal = math.MaxInt32
					for num := range freq {
						if num < minVal {
							minVal = num
						}
					}
				}
			}
			currentLeft++
		}
		
		// Contract from left
		for currentLeft > query.left {
			currentLeft--
			freq[r.nums[currentLeft]]++
			if r.nums[currentLeft] < minVal {
				minVal = r.nums[currentLeft]
			}
		}
		
		results[query.index] = minVal
	}
	
	return results
}

func main() {
	// Test cases
	fmt.Println("=== Testing Mos Algorithm ===")
	
	testCases := []struct {
		nums       []int
		queries    []MosQuery
		description string
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]MosQuery{{0, 2, 0}, {1, 4, 1}, {0, 4, 2}},
			"Basic range sum",
		},
		{
			[]int{1, 1, 2, 2, 3, 3},
			[]MosQuery{{0, 1, 0}, {2, 3, 1}, {4, 5, 2}},
			"Duplicate elements",
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]MosQuery{{0, 4, 0}, {1, 3, 1}, {2, 2, 2}},
			"Descending order",
		},
		{
			[]int{1, 2, 1, 2, 1, 2},
			[]MosQuery{{0, 5, 0}, {1, 4, 1}, {2, 3, 2}},
			"Alternating pattern",
		},
		{
			[]int{10, 20, 30, 40, 50},
			[]MosQuery{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}},
			"Single element queries",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Array: %v\n", tc.nums)
		fmt.Printf("  Queries: %v\n", tc.queries)
		
		rsq := Constructor(tc.nums)
		results := rsq.SumRangeMos(tc.queries)
		fmt.Printf("  Results: %v\n", results)
		
		fmt.Println()
	}
	
	// Test frequency queries
	fmt.Println("=== Frequency Queries Test ===")
	freqNums := []int{1, 2, 3, 2, 1, 2, 3, 2, 1}
	freqQueries := []MosQueryFreq{{0, 4, 2, 0}, {1, 5, 2, 1}, {2, 6, 3, 2}}
	
	rsq := Constructor(freqNums)
	freqResults := rsq.FrequencyQueries(freqQueries)
	fmt.Printf("Frequency queries: %v\n", freqResults)
	
	// Test distinct count queries
	fmt.Println("\n=== Distinct Count Queries Test ===")
	distinctQueries := []MosQueryDistinct{{0, 4, 0}, {1, 5, 1}, {2, 6, 2}}
	distinctResults := rsq.DistinctCountQueries(distinctQueries)
	fmt.Printf("Distinct count queries: %v\n", distinctResults)
	
	// Test mode queries
	fmt.Println("\n=== Mode Queries Test ===")
	modeQueries := []MosQueryMode{{0, 4, 0}, {1, 5, 1}, {2, 6, 2}}
	modeResults := rsq.ModeQueries(modeQueries)
	fmt.Printf("Mode queries: %v\n", modeResults)
	
	// Test min queries
	fmt.Println("\n=== Min Queries Test ===")
	minQueries := []MosQueryMin{{0, 4, 0}, {1, 5, 1}, {2, 6, 2}}
	minResults := rsq.MinQueries(minQueries)
	fmt.Printf("Min queries: %v\n", minResults)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Large array
	largeNums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeNums[i] = i % 1000
	}
	
	// Large queries
	largeQueries := make([]MosQuery, 1000)
	for i := 0; i < 1000; i++ {
		largeQueries[i] = MosQuery{
			left:  i % 5000,
			right: (i % 5000) + 100,
			index: i,
		}
	}
	
	fmt.Printf("Large test with %d elements and %d queries\n", len(largeNums), len(largeQueries))
	
	rsqLarge := Constructor(largeNums)
	start := time.Now()
	results := rsqLarge.SumRangeMos(largeQueries)
	duration := time.Since(start)
	
	fmt.Printf("Large test completed in %v\n", duration)
	fmt.Printf("First 5 results: %v\n", results[:5])
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty array
	emptyRSQ := Constructor([]int{})
	emptyResults := emptyRSQ.SumRangeMos([]MosQuery{})
	fmt.Printf("Empty array: %v\n", emptyResults)
	
	// Single element
	singleRSQ := Constructor([]int{42})
	singleResults := singleRSQ.SumRangeMos([]MosQuery{{0, 0, 0}})
	fmt.Printf("Single element: %v\n", singleResults)
	
	// All same elements
	sameRSQ := Constructor([]int{1, 1, 1, 1, 1})
	sameResults := sameRSQ.DistinctCountQueries([]MosQueryDistinct{{0, 4, 0}})
	fmt.Printf("All same elements: %v\n", sameResults)
	
	// Large range queries
	rangeRSQ := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	rangeResults := rangeRSQ.SumRangeMos([]MosQuery{{0, 9, 0}})
	fmt.Printf("Large range: %v\n", rangeResults)
}
