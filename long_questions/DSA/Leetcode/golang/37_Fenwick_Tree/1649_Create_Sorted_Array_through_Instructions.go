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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fenwick Tree for Dynamic Counting
- **Binary Indexed Tree**: Efficient prefix sum queries and point updates
- **Coordinate Compression**: Handle large value ranges efficiently
- **Dynamic Counting**: Track element frequencies as we process
- **Range Queries**: Count elements less than or greater than current value

## 2. PROBLEM CHARACTERISTICS
- **Sequential Processing**: Process instructions one by one
- **Dynamic Array**: Array grows as we process instructions
- **Cost Calculation**: Cost = min(elements less than, elements greater than)
- **Efficient Counting**: Need fast frequency counting for ranges

## 3. SIMILAR PROBLEMS
- Count of Smaller Numbers After Self (LeetCode 315) - Fenwick tree counting
- Range Sum Query Mutable (LeetCode 307) - Fenwick tree for sums
- Reverse Pairs (LeetCode 493) - Fenwick tree for counting inversions
- Queue Reconstruction by Height (LeetCode 406) - BIT for counting

## 4. KEY OBSERVATIONS
- **Fenwick Tree Natural**: Need efficient prefix sum queries and point updates
- **Dynamic Counting**: Track frequencies of processed elements
- **Cost Calculation**: less = count(< val), greater = count(> val)
- **Coordinate Compression**: Handle large value ranges efficiently

## 5. VARIATIONS & EXTENSIONS
- **Coordinate Compression**: Optimize space for sparse large values
- **Segment Tree**: Alternative data structure for range queries
- **Two Trees**: Separate trees for different counting strategies
- **Frequency Array**: Simple but less efficient approach

## 6. INTERVIEW INSIGHTS
- Always clarify: "Value range constraints? Array size? Need coordinate compression?"
- Edge cases: empty array, single element, all same values
- Time complexity: O(N log M) where M is max value or compressed size
- Space complexity: O(M) for tree, O(N) for compression
- Key insight: Fenwick tree enables O(log M) counting operations

## 7. COMMON MISTAKES
- Not handling coordinate compression for large values
- Wrong Fenwick tree indexing (1-based vs 0-based)
- Incorrect cost calculation (less vs greater)
- Missing modulo operation for large results
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- **Fenwick Tree**: O(N log M) time, O(M) space - standard approach
- **Coordinate Compression**: O(N log N) time, O(N) space - handle large ranges
- **Segment Tree**: O(N log M) time, O(4M) space - alternative approach
- **Frequency Array**: O(NM) time, O(M) space - simple but slow

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a sorted list step by step:**
- Each instruction adds a number to our growing list
- For each new number, we need to know where it fits
- We count how many existing numbers are smaller and larger
- The cost is the minimum of these two counts
- We need an efficient way to maintain counts as we go
- Like maintaining a running frequency distribution

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Instructions array with numbers to insert
2. **Goal**: Create sorted array, calculate total cost
3. **Rules**: Cost = min(count of smaller, count of larger) elements
4. **Output**: Total cost modulo 10^9+7

#### Phase 2: Key Insight Recognition
- **"Dynamic counting natural"** → Need to track element frequencies
- **"Range queries needed"** → Count elements < val and > val
- **"Fenwick tree perfect"** → Efficient prefix sum + point updates
- **"Coordinate compression useful"** → Handle large value ranges

#### Phase 3: Strategy Development
```
Human thought process:
"I need to insert numbers one by one and count smaller/larger.
Naive approach: O(N^2) by scanning previous elements each time.

Fenwick Tree Approach:
1. Use Fenwick tree to maintain frequency counts
2. For each instruction val:
   - less = query(val-1) → count of elements < val
   - greater = i - query(val) → count of elements > val
   - cost = min(less, greater)
   - update(val, 1) → add this element to tree
3. Sum all costs

This gives O(N log M) time!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **Single element**: Cost is 0 (no smaller/larger elements)
- **All same values**: Cost is 0 for all but first element
- **Large values**: Use coordinate compression

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: instructions = [1, 5, 6, 2]

Human thinking:
"Fenwick Tree Approach:
Step 1: Process val=1
- less = query(0) = 0 (no elements < 1)
- greater = 0 - query(1) = 0 (no elements > 1)
- cost = min(0, 0) = 0
- update(1, 1)
Tree now has: freq[1] = 1

Step 2: Process val=5
- less = query(4) = 1 (one element < 5: the 1)
- greater = 1 - query(5) = 0 (no elements > 5)
- cost = min(1, 0) = 0
- update(5, 1)
Tree now has: freq[1] = 1, freq[5] = 1

Step 3: Process val=6
- less = query(5) = 2 (elements < 6: 1, 5)
- greater = 2 - query(6) = 0 (no elements > 6)
- cost = min(2, 0) = 0
- update(6, 1)
Tree now has: freq[1] = 1, freq[5] = 1, freq[6] = 1

Step 4: Process val=2
- less = query(1) = 1 (one element < 2: the 1)
- greater = 3 - query(2) = 2 (elements > 2: 5, 6)
- cost = min(1, 2) = 1
- update(2, 1)

Total cost = 0 + 0 + 0 + 1 = 1 ✓"
```

#### Phase 6: Intuition Validation
- **Why Fenwick tree**: Efficient prefix sum queries and point updates
- **Why O(log M)**: Each query/update traverses O(log M) tree nodes
- **Why coordinate compression**: Reduces space for sparse large values
- **Why dynamic counting**: Build solution incrementally

### Common Human Pitfalls & How to Avoid Them
1. **"Why not simple array?"** → O(N^2) vs O(N log M) with Fenwick tree
2. **"Should I use segment tree?"** → Yes, but Fenwick is simpler for this case
3. **"What about coordinate compression?"** → Essential for large value ranges
4. **"Can I use frequency array?"** → Yes, but inefficient for large ranges
5. **"Why modulo needed?"** → Prevent integer overflow for large results

### Real-World Analogy
**Like managing a dynamic leaderboard in a game:**
- Each player joins with a score (instruction value)
- For each new player, you need to know their rank
- You count how many existing players have lower/higher scores
- The "cost" is like determining if they're in bottom half or top half
- You need efficient updates as players join continuously
- Like maintaining a real-time ranking system

### Human-Readable Pseudocode
```
function createSortedArray(instructions):
    if instructions is empty:
        return 0
    
    # Find maximum value for tree size
    maxVal = max(instructions)
    
    # Initialize Fenwick tree
    tree = array of size maxVal + 1
    
    result = 0
    mod = 10^9 + 7
    
    for i, val in enumerate(instructions):
        # Count elements less than val
        less = query(tree, val - 1)
        
        # Count elements greater than val
        greater = i - query(tree, val)
        
        # Add cost to result
        result = (result + min(less, greater)) % mod
        
        # Update tree with current element
        update(tree, val, 1)
    
    return result

function update(tree, index, delta):
    while index < len(tree):
        tree[index] += delta
        index += index & (-index)  # Move to next node

function query(tree, index):
    sum = 0
    while index > 0:
        sum += tree[index]
        index -= index & (-index)  # Move to parent
    return sum
```

### Execution Visualization

### Example: instructions = [1, 5, 6, 2]
```
Fenwick Tree Process:

Step 1: Process val=1
Tree: [0, 1, 0, 0, 0, 0, 0] (indices 0-6)
less = query(0) = 0
greater = 0 - query(1) = 0
cost = min(0, 0) = 0
Result: 0

Step 2: Process val=5
Update tree[5] = 1
Tree: [0, 1, 0, 0, 0, 1, 0]
less = query(4) = 1 (only the 1)
greater = 1 - query(5) = 0
cost = min(1, 0) = 0
Result: 0

Step 3: Process val=6
Update tree[6] = 1
Tree: [0, 1, 0, 0, 0, 1, 1]
less = query(5) = 2 (1 and 5)
greater = 2 - query(6) = 0
cost = min(2, 0) = 0
Result: 0

Step 4: Process val=2
Update tree[2] = 1
Tree: [0, 1, 1, 0, 0, 1, 1]
less = query(1) = 1 (only the 1)
greater = 3 - query(2) = 2 (5 and 6)
cost = min(1, 2) = 1
Result: 1

Total Cost: 1 ✓
```

### Key Visualization Points:
- **Dynamic Counting**: Tree tracks frequencies of processed elements
- **Prefix Queries**: query(x) gives count of elements ≤ x
- **Range Queries**: less = query(val-1), greater = i - query(val)
- **Point Updates**: update(val, 1) adds current element

### Fenwick Tree Structure:
```
Binary Indexed Tree Structure:
Index: 1  2  3  4  5  6
Tree: [1, 1, 1, 0, 1, 1] (after processing [1,5,6,2])

Query(1) = tree[1] = 1 (count ≤ 1)
Query(2) = tree[2] + tree[1] = 1 + 1 = 2 (count ≤ 2)
Query(5) = tree[5] + tree[4] = 1 + 0 = 1 (count ≤ 5)
```

### Time Complexity Breakdown:
- **Fenwick Tree**: O(N log M) time, O(M) space - standard approach
- **Coordinate Compression**: O(N log N) time, O(N) space - optimized space
- **Segment Tree**: O(N log M) time, O(4M) space - alternative approach
- **Brute Force**: O(N²) time, O(1) space - simple but slow

### Alternative Approaches:

#### 1. Coordinate Compression (O(N log N) time, O(N) space)
```go
func createSortedArrayCompressed(instructions []int) int {
    // Sort unique values and map to indices
    // Use smaller Fenwick tree
    // Handle large value ranges efficiently
    // ... implementation details omitted
}
```
- **Pros**: Optimal space usage for sparse large values
- **Cons**: Additional sorting overhead

#### 2. Segment Tree (O(N log M) time, O(4M) space)
```go
func createSortedArraySegmentTree(instructions []int) int {
    // Build segment tree for range sum queries
    // Same logic as Fenwick tree
    // More memory but flexible for other operations
    // ... implementation details omitted
}
```
- **Pros**: More flexible for different operations
- **Cons**: Higher memory usage, more complex

#### 3. Frequency Array (O(NM) time, O(M) space)
```go
func createSortedArrayFreq(instructions []int) int {
    // Simple frequency counting
    // Linear scan for each query
    // Works for small value ranges only
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Too slow for large value ranges

### Extensions for Interviews:
- **Different Cost Functions**: max(less, greater) or weighted costs
- **Multiple Queries**: Handle batch queries efficiently
- **Range Updates**: Support adding ranges of values
- **2D Extension**: Extend to 2D counting problems
- **Real-world Applications**: Leaderboard systems, ranking algorithms
*/
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
