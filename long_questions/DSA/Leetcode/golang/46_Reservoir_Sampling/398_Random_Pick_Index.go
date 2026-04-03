package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 398. Random Pick Index - Reservoir Sampling
// Time: O(N), Space: O(1)
func pickRandomIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// Reservoir sampling algorithm
	k := 1 // We need to pick 1 element
	reservoir := make([]int, k)
	
	// Fill reservoir with first k elements
	for i := 0; i < k && i < len(nums); i++ {
		reservoir[i] = nums[i]
	}
	
	// Process remaining elements
	for i := k; i < len(nums); i++ {
		// Generate random number between 0 and i
		j := rand.Intn(i + 1)
		if j < k {
			reservoir[j] = nums[i]
		}
	}
	
	return reservoir[0]
}

// Reservoir sampling with multiple picks
func pickRandomIndexMultiple(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return []int{}
	}
	
	reservoir := make([]int, k)
	
	// Fill reservoir with first k elements
	for i := 0; i < k && i < len(nums); i++ {
		reservoir[i] = nums[i]
	}
	
	// Process remaining elements
	for i := k; i < len(nums); i++ {
		j := rand.Intn(i + 1)
		if j < k {
			reservoir[j] = nums[i]
		}
	}
	
	return reservoir
}

// Reservoir sampling with weighted selection
func pickRandomWeightedIndex(nums []int, weights []int) int {
	if len(nums) == 0 || len(nums) != len(weights) {
		return -1
	}
	
	// Calculate cumulative weights
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	
	// Generate random number and find corresponding index
	r := rand.Intn(totalWeight)
	cumulative := 0
	
	for i, weight := range weights {
		cumulative += weight
		if r < cumulative {
			return nums[i]
		}
	}
	
	return nums[len(nums)-1]
}

// Reservoir sampling with streaming data
type ReservoirSampler struct {
	k          int
	reservoir   []int
	count       int
}

func NewReservoirSampler(k int) *ReservoirSampler {
	return &ReservoirSampler{
		k:        k,
		reservoir: make([]int, k),
		count:     0,
	}
}

func (rs *ReservoirSampler) Process(item int) {
	if rs.count < rs.k {
		// Fill reservoir
		rs.reservoir[rs.count] = item
	} else {
		// Random replacement
		j := rand.Intn(rs.count + 1)
		if j < rs.k {
			rs.reservoir[j] = item
		}
	}
	rs.count++
}

func (rs *ReservoirSampler) GetSample() []int {
	if rs.count < rs.k {
		return rs.reservoir[:rs.count]
	}
	return rs.reservoir
}

// Reservoir sampling with deduplication
func pickRandomIndexDedup(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// Use reservoir sampling with set to avoid duplicates
	seen := make(map[int]bool)
	reservoir := make([]int, 1)
	reservoir[0] = nums[0]
	seen[nums[0]] = true
	
	for i := 1; i < len(nums); i++ {
		if !seen[nums[i]] {
			j := rand.Intn(i + 1)
			if j < 1 {
				reservoir[0] = nums[i]
				seen[nums[i]] = true
			}
		}
	}
	
	return reservoir[0]
}

// Reservoir sampling with probability
func pickRandomWithProbability(nums []int, probability float64) int {
	if len(nums) == 0 {
		return -1
	}
	
	if rand.Float64() < probability {
		// Use reservoir sampling
		return pickRandomIndex(nums)
	}
	
	// Return first element
	return nums[0]
}

// Reservoir sampling with time complexity analysis
func pickRandomIndexWithAnalysis(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, 0
	}
	
	start := time.Now()
	result := pickRandomIndex(nums)
	duration := time.Since(start)
	
	// Count operations (simplified)
	operations := len(nums) // Simplified count
	
	return result, operations
}

// Reservoir sampling for large datasets
func pickRandomIndexLarge(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	
	// For very large datasets, we might want to use a different approach
	// But for consistency, we'll use the same algorithm
	return pickRandomIndex(nums)
}

// Reservoir sampling with seed for reproducibility
func pickRandomIndexSeeded(nums []int, seed int64) int {
	if len(nums) == 0 {
		return -1
	}
	
	// Create new random source with seed
	source := rand.NewSource(seed)
	
	k := 1
	reservoir := make([]int, k)
	
	// Fill reservoir
	for i := 0; i < k && i < len(nums); i++ {
		reservoir[i] = nums[i]
	}
	
	// Process remaining elements
	for i := k; i < len(nums); i++ {
		j := source.Intn(i + 1)
		if j < k {
			reservoir[j] = nums[i]
		}
	}
	
	return reservoir[0]
}

// Reservoir sampling with multiple seeds
func pickRandomIndexMultipleSeeds(nums []int, seeds []int64) []int {
	if len(nums) == 0 || len(seeds) == 0 {
		return []int{}
	}
	
	results := make([]int, len(seeds))
	
	for i, seed := range seeds {
		results[i] = pickRandomIndexSeeded(nums, seed)
	}
	
	return results
}

// Reservoir sampling with performance comparison
func compareRandomMethods(nums []int, iterations int) map[string]float64 {
	if len(nums) == 0 {
		return map[string]float64{}
	}
	
	times := make(map[string]time.Duration)
	
	// Test reservoir sampling
	start := time.Now()
	for i := 0; i < iterations; i++ {
		pickRandomIndex(nums)
	}
	times["reservoir"] = time.Since(start)
	
	// Test simple random
	start = time.Now()
	for i := 0; i < iterations; i++ {
		rand.Intn(len(nums))
	}
	times["simple"] = time.Since(start)
	
	// Calculate average times
	averages := make(map[string]float64)
	for method, duration := range times {
		averages[method] = float64(duration.Nanoseconds()) / float64(iterations)
	}
	
	return averages
}

// Reservoir sampling with different k values
func pickRandomIndexVariableK(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return []int{}
	}
	
	reservoir := make([]int, k)
	
	// Fill reservoir with first k elements
	for i := 0; i < k && i < len(nums); i++ {
		reservoir[i] = nums[i]
	}
	
	// Process remaining elements
	for i := k; i < len(nums); i++ {
		j := rand.Intn(i + 1)
		if j < k {
			reservoir[j] = nums[i]
		}
	}
	
	return reservoir
}

func main() {
	// Seed for reproducibility
	rand.Seed(42)
	
	// Test cases
	fmt.Println("=== Testing Random Pick Index - Reservoir Sampling ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{1, 2, 3, 4, 5}, "Standard case"},
		{[]int{10, 20, 30, 40, 50}, "Large numbers"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
		{[]int{1, 1, 2, 2, 3, 3}, "With duplicates"},
		{[]int{-1, -2, -3, -4, -5}, "Negative numbers"},
		{[]int{100, 200, 300, 400, 500}, "Very large numbers"},
		{[]int{0, 1, 2, 3, 4}, "With zero"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Array: %v\n", tc.nums)
		
		result1 := pickRandomIndex(tc.nums)
		result2 := pickRandomIndexMultiple(tc.nums, 3)
		result3 := pickRandomIndexDedup(tc.nums)
		
		fmt.Printf("  Single pick: %d\n", result1)
		fmt.Printf("  Multiple picks: %v\n", result2)
		fmt.Printf("  Dedup pick: %d\n", result3)
		
		fmt.Println()
	}
	
	// Test weighted selection
	fmt.Println("=== Weighted Selection Test ===")
	weights := []int{1, 2, 3, 4, 5}
	weightedResult := pickRandomWeightedIndex([]int{10, 20, 30, 40, 50}, weights)
	fmt.Printf("Weighted result: %d\n", weightedResult)
	
	// Test probability
	fmt.Println("\n=== Probability Test ===")
	probResult := pickRandomWithProbability([]int{1, 2, 3, 4, 5}, 0.5)
	fmt.Printf("Probability (0.5) result: %d\n", probResult)
	
	// Test streaming sampler
	fmt.Println("\n=== Streaming Sampler Test ===")
	streamingNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sampler := NewReservoirSampler(3)
	
	for _, num := range streamingNums {
		sampler.Process(num)
	}
	
	streamResult := sampler.GetSample()
	fmt.Printf("Streaming sampler result: %v\n", streamResult)
	
	// Test performance analysis
	fmt.Println("\n=== Performance Analysis ===")
	largeNums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeNums[i] = i
	}
	
	result, operations := pickRandomIndexWithAnalysis(largeNums)
	fmt.Printf("Large dataset result: %d, operations: %d\n", result, operations)
	
	// Test reproducibility
	fmt.Println("\n=== Reproducibility Test ===")
	seeds := []int64{42, 100, 200}
	reproducibleResults := pickRandomIndexMultipleSeeds([]int{1, 2, 3, 4, 5}, seeds)
	fmt.Printf("Seeds %v: %v\n", seeds, reproducibleResults)
	
	// Test performance comparison
	fmt.Println("\n=== Performance Comparison ===")
	perfNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	averages := compareRandomMethods(perfNums, 1000)
	
	for method, avg := range averages {
		fmt.Printf("%s average: %.2f ns\n", method, avg)
	}
	
	// Test variable k
	fmt.Println("\n=== Variable K Test ===")
	varKResult := pickRandomIndexVariableK([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)
	fmt.Printf("Variable k=4 result: %v\n", varKResult)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Test with very large k
	overKResult := pickRandomIndexMultiple([]int{1, 2, 3}, 5)
	fmt.Printf("k > len(nums): %v\n", overKResult)
	
	// Test with zero k
	zeroKResult := pickRandomIndexMultiple([]int{1, 2, 3}, 0)
	fmt.Printf("k=0: %v\n", zeroKResult)
	
	// Test with negative probability
	negProbResult := pickRandomWithProbability([]int{1, 2, 3}, -0.5)
	fmt.Printf("Negative probability: %d\n", negProbResult)
	
	// Test with single element array
	singleResult := pickRandomIndexMultiple([]int{42}, 1)
	fmt.Printf("Single element: %v\n", singleResult)
	
	// Test consistency
	fmt.Println("\n=== Consistency Test ===")
	consistentNums := []int{1, 2, 3, 4, 5}
	
	// Run multiple times and check distribution
	distribution := make(map[int]int)
	for i := 0; i < 1000; i++ {
		result := pickRandomIndex(consistentNums)
		distribution[result]++
	}
	
	fmt.Printf("Distribution over 1000 runs:\n")
	for num, count := range distribution {
		fmt.Printf("  %d: %d (%.1f%%)\n", num, count, float64(count)/10.0)
	}
}
