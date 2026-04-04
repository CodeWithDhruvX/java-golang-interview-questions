package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 382. Linked List Random Node - Randomized Algorithms
// Time: O(N), Space: O(1)
type ListNode struct {
	Val  int
	Next *ListNode
}

type RandomizedLinkedList struct {
	head *ListNode
}

func Constructor(head *ListNode) RandomizedLinkedList {
	return RandomizedLinkedList{head: head}
}

// Reservoir sampling for random node selection
func (rll *RandomizedLinkedList) GetRandom() int {
	if rll.head == nil {
		return -1
	}
	
	// Reservoir sampling algorithm
	count := 0
	result := rll.head.Val
	
	current := rll.head
	for current != nil {
		count++
		// With probability 1/count, select current node
		if rand.Intn(count) == 0 {
			result = current.Val
		}
		current = current.Next
	}
	
	return result
}

// Monte Carlo algorithm for approximate counting
func (rll *RandomizedLinkedList) ApproximateCount(iterations int) int {
	if rll.head == nil {
		return 0
	}
	
	// Monte Carlo sampling to estimate length
	sampleCount := 0
	
	for i := 0; i < iterations; i++ {
		current := rll.head
		steps := 0
		
		// Random walk to a random position
		for current != nil {
			steps++
			if rand.Float64() < 0.5 {
				break
			}
			current = current.Next
		}
		
		sampleCount += steps
	}
	
	// Estimate total length
	estimatedLength := sampleCount * 2 / iterations
	return estimatedLength
}

// Las Vegas algorithm for finding kth random element
func (rll *RandomizedLinkedList) GetKthRandom(k int) int {
	if rll.head == nil || k <= 0 {
		return -1
	}
	
	// Use reservoir sampling to select k random elements
	reservoir := make([]int, k)
	
	// Fill reservoir with first k elements
	current := rll.head
	count := 0
	
	for current != nil && count < k {
		reservoir[count] = current.Val
		current = current.Next
		count++
	}
	
	// Process remaining elements
	for current != nil {
		count++
		// Randomly replace elements in reservoir
		j := rand.Intn(count)
		if j < k {
			reservoir[j] = current.Val
		}
		current = current.Next
	}
	
	// Return kth element from reservoir
	return reservoir[k-1]
}

// Randomized algorithm for cycle detection
func (rll *RandomizedLinkedList) DetectCycleMonteCarlo(iterations int) bool {
	if rll.head == nil {
		return false
	}
	
	// Monte Carlo approach: sample random pairs of nodes
	for i := 0; i < iterations; i++ {
		// Pick two random positions
		pos1 := rand.Intn(1000) // Assume max length 1000
		pos2 := rand.Intn(1000)
		
		// Find nodes at these positions
		node1 := rll.getNodeAtPosition(pos1)
		node2 := rll.getNodeAtPosition(pos2)
		
		if node1 != nil && node2 != nil && node1 == node2 {
			return true
		}
	}
	
	return false
}

func (rll *RandomizedLinkedList) getNodeAtPosition(pos int) *ListNode {
	current := rll.head
	count := 0
	
	for current != nil && count < pos {
		current = current.Next
		count++
	}
	
	return current
}

// Randomized algorithm for median approximation
func (rll *RandomizedLinkedList) ApproximateMedian(samples int) int {
	if rll.head == nil {
		return -1
	}
	
	// Sample random nodes and compute median of samples
	sampleValues := make([]int, 0)
	
	for i := 0; i < samples; i++ {
		// Get random node using reservoir sampling
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		sampleValues = append(sampleValues, result)
	}
	
	// Sort sample values and return median
	for i := 0; i < len(sampleValues)-1; i++ {
		for j := i + 1; j < len(sampleValues); j++ {
			if sampleValues[i] > sampleValues[j] {
				sampleValues[i], sampleValues[j] = sampleValues[j], sampleValues[i]
			}
		}
	}
	
	return sampleValues[len(sampleValues)/2]
}

// Randomized algorithm for finding mode
func (rll *RandomizedLinkedList) FindModeMonteCarlo(iterations int) int {
	if rll.head == nil {
		return -1
	}
	
	// Monte Carlo approach: sample random nodes
	frequency := make(map[int]int)
	
	for i := 0; i < iterations; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		frequency[result]++
	}
	
	// Find most frequent value
	maxFreq := 0
	mode := -1
	
	for value, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
			mode = value
		}
	}
	
	return mode
}

// Randomized algorithm for finding maximum
func (rll *RandomizedLinkedList) FindMaximumMonteCarlo(iterations int) int {
	if rll.head == nil {
		return -1
	}
	
	// Monte Carlo approach: sample random nodes
	maximum := rll.head.Val
	
	for i := 0; i < iterations; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		if result > maximum {
			maximum = result
		}
	}
	
	return maximum
}

// Randomized algorithm for finding minimum
func (rll *RandomizedLinkedList) FindMinimumMonteCarlo(iterations int) int {
	if rll.head == nil {
		return -1
	}
	
	// Monte Carlo approach: sample random nodes
	minimum := rll.head.Val
	
	for i := 0; i < iterations; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		if result < minimum {
			minimum = result
		}
	}
	
	return minimum
}

// Randomized algorithm for finding average
func (rll *RandomizedLinkedList) ApproximateAverage(samples int) float64 {
	if rll.head == nil {
		return 0.0
	}
	
	// Sample random nodes and compute average
	sum := 0
	
	for i := 0; i < samples; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		sum += result
	}
	
	return float64(sum) / float64(samples)
}

// Randomized algorithm for finding duplicates
func (rll *RandomizedLinkedList) FindDuplicatesMonteCarlo(iterations int) map[int]bool {
	if rll.head == nil {
		return map[int]bool{}
	}
	
	// Monte Carlo approach: sample random nodes
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)
	
	for i := 0; i < iterations; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		if seen[result] {
			duplicates[result] = true
		} else {
			seen[result] = true
		}
	}
	
	return duplicates
}

// Randomized algorithm for finding sum
func (rll *RandomizedLinkedList) ApproximateSum(samples int) int {
	if rll.head == nil {
		return 0
	}
	
	// Sample random nodes and estimate sum
	sampleSum := 0
	
	for i := 0; i < samples; i++ {
		// Get random node
		current := rll.head
		count := 0
		result := rll.head.Val
		
		for current != nil {
			count++
			if rand.Intn(count) == 0 {
				result = current.Val
			}
			current = current.Next
		}
		
		sampleSum += result
	}
	
	// Estimate total sum by scaling
	estimatedLength := rll.ApproximateCount(100)
	return sampleSum * estimatedLength / samples
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Randomized Algorithms for Data Structures
- **Reservoir Sampling**: Uniform random selection from streaming data
- **Monte Carlo Methods**: Probabilistic approximation algorithms
- **Las Vegas Algorithms**: Always correct, randomized runtime
- **Random Walk**: Stochastic traversal of data structures

## 2. PROBLEM CHARACTERISTICS
- **Random Selection**: Choose elements with uniform probability
- **Approximation**: Trade accuracy for efficiency
- **Streaming Data**: Process data without full storage
- **Probabilistic Guarantees: Statistical correctness bounds

## 3. SIMILAR PROBLEMS
- Random Pick Index (LeetCode 398) - Reservoir sampling
- Shuffle an Array (LeetCode 384) - Fisher-Yates shuffle
- Random Point in Non-overlapping Rectangles (LeetCode 497) - Random geometry
- Insert Delete GetRandom O(1) (LeetCode 380) - Random data structure

## 4. KEY OBSERVATIONS
- **Reservoir Sampling**: 1/i probability for ith element ensures uniformity
- **Monte Carlo**: Approximate answers with bounded error probability
- **Las Vegas**: Exact answers with randomized performance
- **Space Efficiency**: O(1) space for streaming algorithms

## 5. VARIATIONS & EXTENSIONS
- **Single Random Selection**: Basic reservoir sampling
- **K Random Elements**: Extended reservoir sampling
- **Approximate Aggregates**: Statistical estimation
- **Randomized Search**: Monte Carlo optimization

## 6. INTERVIEW INSIGHTS
- Always clarify: "Accuracy requirements? Time constraints? Space limits?"
- Edge cases: empty list, single element, large datasets
- Time complexity: O(N) for reservoir sampling, O(k) for k elements
- Space complexity: O(1) for streaming, O(k) for k elements
- Key insight: probability theory enables uniform sampling without storage

## 7. COMMON MISTAKES
- Wrong probability calculations in reservoir sampling
- Confusing Monte Carlo and Las Vegas algorithms
- Not handling edge cases properly
- Incorrect random number generation
- Missing statistical guarantees

## 8. OPTIMIZATION STRATEGIES
- **Reservoir Sampling**: O(N) time, O(1) space - optimal for streaming
- **Array Storage**: O(1) time, O(N) space - fast but memory intensive
- **Hash Map**: O(1) time, O(N) space - for indexed access
- **Approximation**: O(N) time, O(1) space - trade accuracy for speed

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like drawing lottery tickets from a giant bowl:**
- You want to pick a completely random ticket without seeing all tickets
- You go through tickets one by one, deciding whether to keep or replace
- Each new ticket has a specific probability to replace your current choice
- In the end, every ticket had equal chance to be selected
- Like a fair lottery where you can't see all tickets at once

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Linked list with unknown length
2. **Goal**: Select random node with uniform probability
3. **Constraints**: Cannot store entire list, O(1) space
4. **Output**: Random node value

#### Phase 2: Key Insight Recognition
- **"Streaming natural"** → Need algorithm that works without full knowledge
- **"Probability critical"** → Each element must have equal selection chance
- **"Reservoir sampling"** → Classic solution for uniform random selection
- **"1/i probability"** → Mathematical guarantee of uniformity

#### Phase 3: Strategy Development
```
Human thought process:
"I need random node from linked list I can't store.
Brute force: store all nodes in array O(N) space.

Reservoir Sampling Approach:
1. Start with first node as selection
2. For each ith node (starting from 2):
   - With probability 1/i, replace current selection
   - With probability (i-1)/i, keep current selection
3. Each node has exactly 1/n chance of being final selection
4. O(N) time, O(1) space!

This gives perfect uniform distribution!"
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return error value (-1)
- **Single node**: Return that node (only choice)
- **Large lists**: Algorithm scales perfectly
- **Random seed**: Ensure reproducibility for testing

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: List [1, 2, 3, 4, 5]

Human thinking:
"Reservoir Sampling Process:
Step 1: Start with node 1
Selection = 1, Count = 1

Step 2: Process node 2
Count = 2
With probability 1/2: replace selection with 2
With probability 1/2: keep selection as 1
Assume we keep 1

Step 3: Process node 3
Count = 3
With probability 1/3: replace selection with 3
With probability 2/3: keep current selection
Assume we keep 1

Step 4: Process node 4
Count = 4
With probability 1/4: replace with 4
With probability 3/4: keep current
Assume we replace with 4

Step 5: Process node 5
Count = 5
With probability 1/5: replace with 5
With probability 4/5: keep 4
Assume we keep 4

Final selection: 4
Each node had exactly 1/5 chance! ✓"
```

#### Phase 6: Intuition Validation
- **Why 1/i probability**: Mathematical proof shows uniform distribution
- **Why O(1) space**: Only store current selection and count
- **Why O(N) time**: Single pass through list
- **Why uniform**: Each node's selection probability = 1/n

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count and pick random index?"** → Requires O(N) space
2. **"Should I use different probabilities?"** → 1/i is mathematically optimal
3. **"What about multiple selections?"** → Use reservoir sampling with k elements
4. **"Can I use true randomness?"** → Pseudorandom is sufficient
5. **"Why not store in array?"** → Violates O(1) space constraint

### Real-World Analogy
**Like a restaurant manager hiring from a long line of applicants:**
- You can't interview everyone, need to sample randomly
- You interview people one by one as they arrive
- Each new applicant has a chance to replace your current favorite
- By the end, every applicant had equal chance to be selected
- Like fair hiring when you can't see all candidates at once

### Human-Readable Pseudocode
```
function getRandomNode(linkedList):
    if linkedList.head == null:
        return -1
    
    selection = linkedList.head.val
    count = 1
    current = linkedList.head.next
    
    while current != null:
        count += 1
        
        # With probability 1/count, replace selection
        if randomInt(0, count-1) == 0:
            selection = current.val
        
        current = current.next
    
    return selection

# For k random elements
function getKRandomNodes(linkedList, k):
    if linkedList.head == null:
        return []
    
    # Fill reservoir with first k elements
    reservoir = []
    current = linkedList.head
    count = 0
    
    while current != null and count < k:
        reservoir[count] = current.val
        current = current.next
        count += 1
    
    # Process remaining elements
    while current != null:
        count += 1
        
        # Randomly replace elements in reservoir
        j = randomInt(0, count-1)
        if j < k:
            reservoir[j] = current.val
        
        current = current.next
    
    return reservoir
```

### Execution Visualization

### Example: List [1, 2, 3, 4, 5]
```
Reservoir Sampling Process:

Step 1: Initialize
Selection = 1, Count = 1

Step 2: Process node 2 (value = 2)
Count = 2
Random number in [0,1]: let's say 1
Since 1 != 0, keep selection = 1
Probability of keeping: 1/2, replacing: 1/2

Step 3: Process node 3 (value = 3)
Count = 3
Random number in [0,2]: let's say 0
Since 0 == 0, replace selection = 3
Probability of keeping: 2/3, replacing: 1/3

Step 4: Process node 4 (value = 4)
Count = 4
Random number in [0,3]: let's say 2
Since 2 != 0, keep selection = 3
Probability of keeping: 3/4, replacing: 1/4

Step 5: Process node 5 (value = 5)
Count = 5
Random number in [0,4]: let's say 1
Since 1 != 0, keep selection = 3
Probability of keeping: 4/5, replacing: 1/5

Final selection: 3

Each node's selection probability:
Node 1: (1/2) × (2/3) × (3/4) × (4/5) = 1/5 ✓
Node 2: (1/2) × (2/3) × (3/4) × (4/5) = 1/5 ✓
Node 3: (1/3) × (3/4) × (4/5) = 1/5 ✓
Node 4: (1/4) × (4/5) = 1/5 ✓
Node 5: (1/5) = 1/5 ✓
```

### Key Visualization Points:
- **Probability Updates**: Each step updates selection probability
- **Uniform Distribution**: Final result is truly uniform
- **Space Efficiency**: Only store current selection
- **Mathematical Guarantee**: Proven uniform distribution

### Monte Carlo vs Las Vegas Visualization:
```
Monte Carlo (Approximate):
- Always fast, sometimes wrong
- Example: Approximate counting with sampling
- Trade accuracy for speed

Las Vegas (Exact):
- Sometimes slow, always correct
- Example: Reservoir sampling for exact random selection
- Trade time for correctness
```

### Time Complexity Breakdown:
- **Reservoir Sampling**: O(N) time, O(1) space - optimal for streaming
- **Array Storage**: O(1) time, O(N) space - fast but memory intensive
- **Hash Map**: O(1) time, O(N) space - for indexed access
- **Approximation**: O(N) time, O(1) space - trade accuracy for speed

### Alternative Approaches:

#### 1. Array Storage (O(1) time, O(N) space)
```go
func getRandomNodeArray(head *ListNode) int {
    // Store all nodes in array
    nodes := []int{}
    current := head
    for current != nil {
        nodes = append(nodes, current.Val)
        current = current.Next
    }
    
    // Pick random index
    return nodes[rand.Intn(len(nodes))]
}
```
- **Pros**: O(1) time for selection, simple implementation
- **Cons**: Requires O(N) space, not suitable for streaming

#### 2. Two-Pass Algorithm (O(N) time, O(1) space)
```go
func getRandomNodeTwoPass(head *ListNode) int {
    // First pass: count nodes
    count := 0
    current := head
    for current != nil {
        count++
        current = current.Next
    }
    
    // Second pass: select random node
    randomIndex := rand.Intn(count)
    current = head
    for i := 0; i < randomIndex; i++ {
        current = current.Next
    }
    
    return current.Val
}
```
- **Pros**: O(1) space, truly uniform
- **Cons**: Two passes through list, not streaming

#### 3. Approximate Sampling (O(N) time, O(1) space)
```go
func getRandomNodeApproximate(head *ListNode, samples int) int {
    // Sample k nodes and pick randomly from sample
    // Not perfectly uniform but good approximation
    // Trade accuracy for simplicity
}
```
- **Pros**: Simple, fast
- **Cons**: Not perfectly uniform

### Extensions for Interviews:
- **Weighted Random Selection**: Different probabilities for elements
- **Streaming Variants**: Handle infinite streams
- **Multiple Selections**: Select k random elements
- **Reservoir Variants**: Different sampling strategies
- **Real-world Applications**: Database sampling, load balancing, A/B testing
*/
func main() {
	// Seed for reproducibility
	rand.Seed(42)
	
	// Test cases
	fmt.Println("=== Testing Randomized Algorithms ===")
	
	// Helper function to create linked list
	createLinkedList := func(values []int) *ListNode {
		if len(values) == 0 {
			return nil
		}
		
		head := &ListNode{Val: values[0]}
		current := head
		
		for i := 1; i < len(values); i++ {
			current.Next = &ListNode{Val: values[i]}
			current = current.Next
		}
		
		return head
	}
	
	testCases := []struct {
		values     []int
		description string
	}{
		{[]int{1, 2, 3, 4, 5}, "Standard case"},
		{[]int{10, 20, 30, 40, 50}, "Large numbers"},
		{[]int{1}, "Single node"},
		{[]int{}, "Empty list"},
		{[]int{1, 1, 2, 2, 3, 3}, "With duplicates"},
		{[]int{-1, -2, -3, -4, -5}, "Negative numbers"},
		{[]int{100, 200, 300, 400, 500}, "Very large numbers"},
		{[]int{0, 1, 2, 3, 4}, "With zero"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Values: %v\n", tc.values)
		
		head := createLinkedList(tc.values)
		rll := Constructor(head)
		
		// Test random selection
		randomResult := rll.GetRandom()
		fmt.Printf("  Random node: %d\n", randomResult)
		
		// Test approximate counting
		approxCount := rll.ApproximateCount(100)
		fmt.Printf("  Approximate count: %d\n", approxCount)
		
		// Test kth random
		if len(tc.values) >= 3 {
			kthRandom := rll.GetKthRandom(3)
			fmt.Printf("  3rd random: %d\n", kthRandom)
		}
		
		// Test cycle detection
		hasCycle := rll.DetectCycleMonteCarlo(100)
		fmt.Printf("  Has cycle: %t\n", hasCycle)
		
		// Test approximate median
		approxMedian := rll.ApproximateMedian(50)
		fmt.Printf("  Approximate median: %d\n", approxMedian)
		
		// Test mode finding
		mode := rll.FindModeMonteCarlo(100)
		fmt.Printf("  Mode: %d\n", mode)
		
		// Test maximum finding
		maximum := rll.FindMaximumMonteCarlo(100)
		fmt.Printf("  Maximum: %d\n", maximum)
		
		// Test minimum finding
		minimum := rll.FindMinimumMonteCarlo(100)
		fmt.Printf("  Minimum: %d\n", minimum)
		
		// Test average approximation
		average := rll.ApproximateAverage(50)
		fmt.Printf("  Approximate average: %.2f\n", average)
		
		// Test duplicate finding
		duplicates := rll.FindDuplicatesMonteCarlo(100)
		fmt.Printf("  Duplicates: %v\n", duplicates)
		
		// Test sum approximation
		approxSum := rll.ApproximateSum(50)
		fmt.Printf("  Approximate sum: %d\n", approxSum)
		
		fmt.Println()
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	// Create large linked list
	largeValues := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeValues[i] = i % 1000
	}
	
	largeHead := createLinkedList(largeValues)
	largeRLL := Constructor(largeHead)
	
	fmt.Printf("Large list with %d nodes\n", len(largeValues))
	
	start := time.Now()
	randomResult := largeRLL.GetRandom()
	duration := time.Since(start)
	
	fmt.Printf("Random selection: %d, Time: %v\n", randomResult, duration)
	
	start = time.Now()
	approxCount := largeRLL.ApproximateCount(1000)
	duration = time.Since(start)
	
	fmt.Printf("Approximate count: %d, Time: %v\n", approxCount, duration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Test with nil head
	nilRLL := Constructor(nil)
	fmt.Printf("Nil list random: %d\n", nilRLL.GetRandom())
	fmt.Printf("Nil list count: %d\n", nilRLL.ApproximateCount(100))
	
	// Test with single node
	singleHead := createLinkedList([]int{42})
	singleRLL := Constructor(singleHead)
	fmt.Printf("Single node random: %d\n", singleRLL.GetRandom())
	fmt.Printf("Single node count: %d\n", singleRLL.ApproximateCount(100))
	
	// Test consistency
	fmt.Println("\n=== Consistency Test ===")
	
	// Test multiple random selections
	consistencyHead := createLinkedList([]int{1, 2, 3, 4, 5})
	consistencyRLL := Constructor(consistencyHead)
	
	fmt.Printf("Multiple random selections: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", consistencyRLL.GetRandom())
	}
	fmt.Println()
	
	// Test approximation accuracy
	fmt.Println("\n=== Approximation Accuracy Test ===")
	
	accuracyHead := createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	accuracyRLL := Constructor(accuracyHead)
	
	fmt.Printf("Actual count: %d\n", 10)
	for samples := 10; samples <= 100; samples += 10 {
		approx := accuracyRLL.ApproximateCount(samples)
		fmt.Printf("Samples %d: Approximate count: %d, Error: %d\n", samples, approx, approx-10)
	}
	
	// Test Monte Carlo vs Las Vegas
	fmt.Println("\n=== Monte Carlo vs Las Vegas Test ===")
	
	mcHead := createLinkedList([]int{1, 2, 3, 4, 5})
	mcRLL := Constructor(mcHead)
	
	fmt.Printf("Monte Carlo cycle detection: %t\n", mcRLL.DetectCycleMonteCarlo(100))
	fmt.Printf("Las Vegas random selection: %d\n", mcRLL.GetRandom())
}
