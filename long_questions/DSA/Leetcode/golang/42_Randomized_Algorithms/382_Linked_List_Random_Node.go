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
