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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Mo's Algorithm for Offline Range Queries
- **Offline Processing**: Sort queries to minimize range adjustments
- **Block Ordering**: Sort by block number with alternating right ordering
- **Sliding Window**: Maintain current range with add/remove operations
- **Efficient Updates**: O(1) add/remove operations for most queries

## 2. PROBLEM CHARACTERISTICS
- **Multiple Range Queries**: Process many range queries simultaneously
- **Offline Processing**: All queries known in advance
- **Range Adjustments**: Move between queries with minimal changes
- **Flexible Operations**: Support various range operations (sum, frequency, mode)

## 3. SIMILAR PROBLEMS
- Range Sum Query Mutable (LeetCode 307) - Online range queries
- D-Query (SPOJ) - Distinct elements in range
- Powerful Array (Codeforces) - Range mode queries
- K-Query (SPOJ) - Range counting with offline processing

## 4. KEY OBSERVATIONS
- **Query Ordering**: Optimal ordering reduces total adjustments
- **Block Size**: √N block size balances sorting and adjustment costs
- **Alternating Ordering**: Even blocks sort right ascending, odd blocks descending
- **Sliding Window**: Maintain current range state efficiently

## 5. VARIATIONS & EXTENSIONS
- **Frequency Queries**: Count occurrences of specific values
- **Distinct Count**: Track number of distinct elements
- **Mode Queries**: Find most frequent element in range
- **Min/Max Queries**: Track minimum/maximum values in range

## 6. INTERVIEW INSIGHTS
- Always clarify: "Number of queries? Query types? Need online processing?"
- Edge cases: empty array, single element, all same values
- Time complexity: O((N+Q)√N) for Q queries on N elements
- Space complexity: O(N+Q) for array and query storage
- Key insight: offline processing enables efficient range adjustments

## 7. COMMON MISTAKES
- Wrong block size calculation (should be √N)
- Incorrect query ordering (missing alternating pattern)
- Inefficient add/remove operations
- Not handling edge cases in range adjustments
- Missing modulo operations for large results

## 8. OPTIMIZATION STRATEGIES
- **Standard Mo's**: O((N+Q)√N) time, O(N+Q) space - basic approach
- **Hilbert Order**: O((N+Q)√N) time, O(N+Q) space - better locality
- **Block Adjustments**: O((N+Q)√N) time, O(N+Q) space - optimized blocks
- **Parallel Processing**: O((N+Q)√N/P) time, O(N+Q) space - multi-threaded

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like answering many questions about a book efficiently:**
- Each question asks about a specific chapter range
- Instead of re-reading for each question, you slide your reading window
- You order questions to minimize how much you need to move the window
- You maintain notes about what you've currently read
- Like a researcher efficiently answering many related questions

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of elements, multiple range queries
2. **Goal**: Answer all range queries efficiently
3. **Rules**: All queries known in advance (offline)
4. **Output**: Results for all queries in original order

#### Phase 2: Key Insight Recognition
- **"Offline processing possible"** → All queries known beforehand
- **"Range adjustments expensive"** → Want to minimize movement between ranges
- **"Sliding window natural"** → Maintain current range state
- **"Query ordering critical"** → Optimize order to reduce total movement

#### Phase 3: Strategy Development
```
Human thought process:
"I need to answer many range queries.
Individual processing would be O(N*Q) - too slow.

Mo's Algorithm Approach:
1. Sort queries by block number (√N blocks)
2. Within each block, sort right endpoint (alternating direction)
3. Maintain sliding window [L,R] with current answer
4. For each query, adjust window and update answer
5. Total movement minimized by ordering

This gives O((N+Q)√N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty results
- **Single element**: Trivial range adjustments
- **All same values**: Optimized frequency tracking
- **Large ranges**: Efficient block size calculation

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [1,2,3,4,5], queries = [0,2], [1,4], [0,4]

Human thinking:
"Mo's Algorithm Approach:
Step 1: Calculate block size = √5 ≈ 2
Step 2: Sort queries:
- Query [0,2]: block 0, right 2
- Query [0,4]: block 0, right 4  
- Query [1,4]: block 1, right 4
Order: [0,2], [0,4], [1,4] (block 0: ascending, block 1: ascending)

Step 3: Process queries:
Initialize: L=0, R=-1, sum=0

Query [0,2]:
- Expand R to 2: add nums[0]=1, nums[1]=2, nums[2]=3
- sum = 6
- Answer[0] = 6

Query [0,4]:
- Expand R to 4: add nums[3]=4, nums[4]=5  
- sum = 15
- Answer[2] = 15

Query [1,4]:
- Contract L to 1: remove nums[0]=1
- sum = 14
- Answer[1] = 14

Final answers: [6, 14, 15] ✓"
```

#### Phase 6: Intuition Validation
- **Why ordering matters**: Minimizes total range adjustments
- **Why √N blocks**: Balances sorting cost vs adjustment cost
- **Why alternating direction**: Improves cache locality
- **Why O((N+Q)√N)**: Each element moved O(√N) times total

### Common Human Pitfalls & How to Avoid Them
1. **"Why not process queries individually?"** → O(N*Q) vs O((N+Q)√N)
2. **"Should I use segment tree?"** → Online vs offline processing
3. **"What about block size?"** → √N is optimal for most cases
4. **"Can I use different ordering?"** → Hilbert order for better locality
5. **"Why alternating direction?"** → Reduces right pointer movement

### Real-World Analogy
**Like a librarian answering many book-related questions:**
- Each question asks about a range of pages
- Instead of re-reading for each question, you slide bookmarks
- You order questions to minimize page turning
- You keep notes about what you've currently read
- Like efficiently answering research questions about a document

### Human-Readable Pseudocode
```
function mosAlgorithm(nums, queries):
    n = length(nums)
    blockSize = sqrt(n)
    
    # Sort queries using Mo's ordering
    for query in queries:
        query.block = query.left / blockSize
    
    sort(queries, by: block, then by right with alternating direction)
    
    # Initialize sliding window
    currentLeft = 0
    currentRight = -1
    currentAnswer = 0
    
    # Process queries
    for query in queries:
        # Expand to right
        while currentRight < query.right:
            currentRight += 1
            currentAnswer += nums[currentRight]
        
        # Contract from right
        while currentRight > query.right:
            currentAnswer -= nums[currentRight]
            currentRight -= 1
        
        # Expand to left
        while currentLeft < query.left:
            currentAnswer -= nums[currentLeft]
            currentLeft += 1
        
        # Contract from left
        while currentLeft > query.left:
            currentLeft -= 1
            currentAnswer += nums[currentLeft]
        
        # Store answer
        query.result = currentAnswer
    
    return results in original order
```

### Execution Visualization

### Example: nums = [1,2,3,4,5], queries = [0,2], [1,4], [0,4]
```
Mo's Algorithm Process:

Step 1: Block size = √5 ≈ 2
Query blocks: [0,2]→block 0, [1,4]→block 1, [0,4]→block 0

Step 2: Sort queries
Block 0 (even): [0,2], [0,4] (ascending right)
Block 1 (odd): [1,4]
Final order: [0,2], [0,4], [1,4]

Step 3: Process queries
Initialize: L=0, R=-1, sum=0

Query [0,2]:
- R expands: 0→1→2, sum = 1+2+3 = 6
- Answer[0] = 6

Query [0,4]:  
- R expands: 2→3→4, sum = 6+4+5 = 15
- Answer[2] = 15

Query [1,4]:
- L contracts: 0→1, sum = 15-1 = 14
- Answer[1] = 14

Results: [6, 14, 15] ✓
```

### Key Visualization Points:
- **Query Ordering**: Minimizes total pointer movement
- **Sliding Window**: Efficiently maintain current range
- **Block Sorting**: √N blocks balance costs
- **Alternating Direction**: Improves cache locality

### Query Ordering Visualization:
```
Array indices: 0 1 2 3 4 5 6 7 8 9
Block size: 3 (√10 ≈ 3)

Blocks: [0,1,2] [3,4,5] [6,7,8] [9]

Query ordering:
Block 0 (even): sort right ascending
Block 1 (odd): sort right descending  
Block 2 (even): sort right ascending
etc.
```

### Time Complexity Breakdown:
- **Standard Mo's**: O((N+Q)√N) time, O(N+Q) space - basic approach
- **Hilbert Order**: O((N+Q)√N) time, O(N+Q) space - better locality
- **Optimized Blocks**: O((N+Q)√N) time, O(N+Q) space - tuned block size
- **Parallel**: O((N+Q)√N/P) time, O(N+Q) space - multi-threaded

### Alternative Approaches:

#### 1. Segment Tree (O(N log N + Q log N) time, O(N) space)
```go
func rangeSumSegmentTree(nums []int, queries []Query) []int {
    // Build segment tree for range sum queries
    // Answer each query independently
    // Online processing capability
    // ... implementation details omitted
}
```
- **Pros**: Online processing, log N per query
- **Cons**: More complex, slower for many queries

#### 2. Prefix Sum Array (O(N + Q) time, O(N) space)
```go
func rangeSumPrefix(nums []int, queries []Query) []int {
    // Build prefix sum array
    // Each query: prefix[right+1] - prefix[left]
    // Only for immutable arrays
    // ... implementation details omitted
}
```
- **Pros**: O(1) per query, simple implementation
- **Cons**: No updates allowed, immutable only

#### 3. Square Root Decomposition (O(N + Q√N) time, O(N) space)
```go
func rangeSumSqrt(nums []int, queries []Query) []int {
    // Precompute block sums
    // Each query: O(√N) block traversal
    // Similar to Mo's but simpler
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand, good performance
- **Cons**: Slower than Mo's for many queries

### Extensions for Interviews:
- **Different Operations**: Mode, median, distinct count queries
- **Update Support**: Handle array updates with Mo's algorithm variants
- **2D Queries**: Extend to 2D range queries
- **Hilbert Order**: Better query ordering for complex cases
- **Real-world Applications**: Database range queries, analytics processing
*/
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
