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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Reservoir Sampling for Random Selection
- **Uniform Sampling**: Equal probability for all elements in unknown size datasets
- **Streaming Data**: Process elements without knowing total count in advance
- **Memory Efficiency**: O(k) space for selecting k elements from N elements
- **Probability Theory**: Mathematical guarantee of uniform selection

## 2. PROBLEM CHARACTERISTICS
- **Random Selection**: Pick random element from array with equal probability
- **Unknown Size**: Algorithm works without knowing total elements beforehand
- **Streaming Capability**: Can process data as it arrives
- **Memory Constraints**: Limited space for storing samples

## 3. SIMILAR PROBLEMS
- Linked List Random Node (LeetCode 382) - Same reservoir sampling
- Random Pick Index (LeetCode 398) - Same problem
- Random Pick with Weight - Weighted random selection
- Monte Carlo Methods - Random sampling for estimation

## 4. KEY OBSERVATIONS
- **Uniform Probability**: Each element has 1/N chance of being selected
- **Streaming Natural**: Perfect for unknown or infinite data streams
- **Mathematical Proof**: Probability of selection remains constant
- **Memory Efficiency**: O(k) space regardless of total elements

## 5. VARIATIONS & EXTENSIONS
- **Single Selection**: Pick one random element (k=1)
- **Multiple Selection**: Pick k random elements
- **Weighted Selection**: Probability proportional to weights
- **Streaming Processing**: Handle data streams efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size? Streaming data? Weighted selection? Multiple picks?"
- Edge cases: empty array, single element, duplicates
- Time complexity: O(N) time, O(k) space
- Space complexity: O(k) for k samples
- Key insight: perfect for unknown size datasets and streaming data

## 7. COMMON MISTAKES
- Wrong probability calculation (should be 1/(i+1) for i-th element)
- Missing edge cases for empty/single element arrays
- Incorrect reservoir initialization
- Wrong random number range (0 to i, not 0 to N-1)
- Not handling streaming data properly

## 8. OPTIMIZATION STRATEGIES
- **Standard Reservoir**: O(N) time, O(k) space - optimal
- **Weighted Reservoir**: O(N log W) time, O(k) space - for weighted selection
- **Streaming**: O(N) time, O(k) space - perfect for streams
- **Multiple Seeds**: O(N) time, O(k) space - for reproducible results

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like drawing lottery tickets from a giant pile:**
- You have a huge pile of lottery tickets (unknown total)
- You want to draw k winning tickets with equal probability
- You can't count all tickets beforehand
- You maintain a small sample of k tickets
- For each new ticket, you decide whether to replace one in your sample
- Like a lottery draw where each ticket has equal chance regardless of pile size

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of elements, need random selection
2. **Goal**: Pick random element with equal probability
3. **Constraints**: Should work for unknown size, streaming data
4. **Output**: Randomly selected element(s)

#### Phase 2: Key Insight Recognition
- **"Uniform probability needed"** → Each element must have equal chance
- **"Unknown size natural"** → Perfect for reservoir sampling
- **"Streaming capability"** → Process elements as they arrive
- **"Memory efficient"** → Only store k elements regardless of total

#### Phase 3: Strategy Development
```
Human thought process:
"I need to pick random element with equal probability.
Brute force: count all elements, then random index O(N) space.

Reservoir Sampling Approach:
1. Fill reservoir with first k elements
2. For each new element i:
   - Generate random number j between 0 and i
   - If j < k, replace reservoir[j] with new element
3. Each element has k/(i+1) chance to be in reservoir
4. Final reservoir contains uniformly random sample

This gives O(N) time, O(k) space!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 or empty slice
- **Single element**: Return that element (only choice)
- **k > N**: Return all elements or handle appropriately
- **Streaming data**: Process elements as they arrive

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [1, 2, 3, 4, 5], k = 2

Human thinking:
"Reservoir Sampling Process:
Step 1: Initialize reservoir with first 2 elements
reservoir = [1, 2]

Step 2: Process element 3 (i=2)
Generate j = rand.Intn(3) → j ∈ [0,1,2]
If j < 2: replace reservoir[j] with 3
Probability: 2/3 chance to replace

Step 3: Process element 4 (i=3)
Generate j = rand.Intn(4) → j ∈ [0,1,2,3]
If j < 2: replace reservoir[j] with 4
Probability: 2/4 = 1/2 chance to replace

Step 4: Process element 5 (i=4)
Generate j = rand.Intn(5) → j ∈ [0,1,2,3,4]
If j < 2: replace reservoir[j] with 5
Probability: 2/5 chance to replace

Final: reservoir contains 2 uniformly random elements ✓"
```

#### Phase 6: Intuition Validation
- **Why uniform probability**: Each element has k/(i+1) chance when processed
- **Why streaming works**: No need to know total size in advance
- **Why O(k) space**: Only store sample, not entire dataset
- **Why mathematical proof**: Probability calculations confirm uniformity

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count and pick?"** → Requires O(N) space, can't stream
2. **"Should I use random shuffle?"** → Requires O(N) time and space
3. **"What about weighted selection?"** → Different algorithm, use cumulative weights
4. **"Can I pick multiple elements?"** → Yes, set k > 1
5. **"Why random range 0 to i?"** → Ensures uniform probability for i-th element

### Real-World Analogy
**Like quality control sampling in a factory:**
- You have a production line with unknown total items
- You want to sample k items for quality testing
- You can't stop the line to count all items
- You maintain a sample box of k items
- For each new item, you decide whether to swap it into your sample
- Like a quality inspector sampling products from a continuous production line

### Human-Readable Pseudocode
```
function reservoirSampling(stream, k):
    reservoir = []
    
    # Process first k elements
    for i from 0 to k-1:
        reservoir[i] = stream[i]
    
    # Process remaining elements
    for i from k to n-1:
        j = random number from 0 to i
        if j < k:
            reservoir[j] = stream[i]
    
    return reservoir
```

### Execution Visualization

### Example: nums = [1, 2, 3, 4, 5], k = 1
```
Initial: reservoir = [1]

Process element 2 (i=1):
j = rand.Intn(2) → j ∈ [0,1]
If j < 1: replace reservoir[0] with 2
Probability: 1/2 chance to replace

Process element 3 (i=2):
j = rand.Intn(3) → j ∈ [0,1,2]
If j < 1: replace reservoir[0] with 3
Probability: 1/3 chance to replace

Process element 4 (i=3):
j = rand.Intn(4) → j ∈ [0,1,2,3]
If j < 1: replace reservoir[0] with 4
Probability: 1/4 chance to replace

Process element 5 (i=4):
j = rand.Intn(5) → j ∈ [0,1,2,3,4]
If j < 1: replace reservoir[0] with 5
Probability: 1/5 chance to replace

Final: reservoir[0] has 1/5 chance of being any element ✓
```

### Key Visualization Points:
- **Reservoir Maintenance**: Always keep k elements in sample
- **Random Replacement**: Each new element has chance to replace
- **Uniform Probability**: Equal chance for all elements
- **Streaming Capability**: Process elements without knowing total

### Probability Calculation Visualization:
```
For element i (0-indexed):
- Probability to be selected: k/(i+1)
- Probability to survive all future selections: ∏(j=i+1 to n-1) (1 - k/(j+1))
- Final probability: k/n (uniform for all elements)
```

### Time Complexity Breakdown:
- **Standard Reservoir**: O(N) time, O(k) space - optimal
- **Weighted Reservoir**: O(N log W) time, O(k) space - for weighted selection
- **Streaming**: O(N) time, O(k) space - perfect for streams
- **Multiple Selections**: O(N) time, O(k) space - for k > 1

### Alternative Approaches:

#### 1. Simple Random Index (O(1) time, O(1) space)
```go
func pickRandomIndexSimple(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    return nums[rand.Intn(len(nums))]
}
```
- **Pros**: Simple, fast for known arrays
- **Cons**: Can't handle streaming, requires knowing size

#### 2. Weighted Random Selection (O(N) time, O(1) space)
```go
func pickRandomWeightedIndex(nums []int, weights []int) int {
    totalWeight := 0
    for _, w := range weights {
        totalWeight += w
    }
    
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
```
- **Pros**: Handles weighted selection naturally
- **Cons**: Different probability distribution

#### 3. Fisher-Yates Shuffle (O(N) time, O(1) space)
```go
func pickRandomIndexShuffle(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    
    // Shuffle first element
    j := rand.Intn(len(nums))
    nums[0], nums[j] = nums[j], nums[0]
    
    return nums[0]
}
```
- **Pros**: True random permutation
- **Cons**: Modifies original array, more complex

### Extensions for Interviews:
- **Streaming Data**: Handle infinite or unknown size data streams
- **Weighted Sampling**: Probability proportional to element weights
- **Multiple Samples**: Pick k elements without replacement
- **Reproducible Results**: Use seeds for consistent random selection
- **Real-world Applications**: Quality control, statistical sampling, A/B testing
*/
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
