package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 380. Insert Delete GetRandom O(1)
// Time: O(1) for all operations, Space: O(N)
type RandomizedSet struct {
	nums      []int
	numToIdx  map[int]int
}

// Constructor initializes the data structure
func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{
		nums:     []int{},
		numToIdx: make(map[int]int),
	}
}

// Insert inserts a value to the set. Returns true if the set did not contain the specified element.
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.numToIdx[val]; exists {
		return false
	}
	
	this.numToIdx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

// Remove removes a value from the set. Returns true if the set contained the specified element.
func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.numToIdx[val]; !exists {
		return false
	}
	
	// Get index of element to remove
	idx := this.numToIdx[val]
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	// Move last element to the position of element to remove
	this.nums[idx] = lastVal
	this.numToIdx[lastVal] = idx
	
	// Remove last element
	this.nums = this.nums[:lastIdx]
	delete(this.numToIdx, val)
	
	return true
}

// GetRandom returns a random element from the current set.
func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 0 {
		return -1 // Or handle error appropriately
	}
	
	rand.Seed(time.Now().UnixNano())
	randomIdx := rand.Intn(len(this.nums))
	return this.nums[randomIdx]
}

// Alternative implementation with more detailed tracking
type RandomizedSetOptimized struct {
	nums     []int
	numToIdx map[int]int
	randGen  *rand.Rand
}

func ConstructorRandomizedSetOptimized() RandomizedSetOptimized {
	source := rand.NewSource(time.Now().UnixNano())
	return RandomizedSetOptimized{
		nums:     []int{},
		numToIdx: make(map[int]int),
		randGen:  rand.New(source),
	}
}

func (this *RandomizedSetOptimized) Insert(val int) bool {
	if _, exists := this.numToIdx[val]; exists {
		return false
	}
	
	this.numToIdx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSetOptimized) Remove(val int) bool {
	if _, exists := this.numToIdx[val]; !exists {
		return false
	}
	
	idx := this.numToIdx[val]
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	this.nums[idx] = lastVal
	this.numToIdx[lastVal] = idx
	
	this.nums = this.nums[:lastIdx]
	delete(this.numToIdx, val)
	
	return true
}

func (this *RandomizedSetOptimized) GetRandom() int {
	if len(this.nums) == 0 {
		return -1
	}
	
	randomIdx := this.randGen.Intn(len(this.nums))
	return this.nums[randomIdx]
}

// Version with statistics
type RandomizedSetWithStats struct {
	nums          []int
	numToIdx      map[int]int
	insertCount   int
	removeCount   int
	getRandomCount int
	randGen       *rand.Rand
}

func ConstructorRandomizedSetWithStats() RandomizedSetWithStats {
	source := rand.NewSource(time.Now().UnixNano())
	return RandomizedSetWithStats{
		nums:     []int{},
		numToIdx: make(map[int]int),
		randGen:  rand.New(source),
	}
}

func (this *RandomizedSetWithStats) Insert(val int) bool {
	this.insertCount++
	
	if _, exists := this.numToIdx[val]; exists {
		return false
	}
	
	this.numToIdx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSetWithStats) Remove(val int) bool {
	this.removeCount++
	
	if _, exists := this.numToIdx[val]; !exists {
		return false
	}
	
	idx := this.numToIdx[val]
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	this.nums[idx] = lastVal
	this.numToIdx[lastVal] = idx
	
	this.nums = this.nums[:lastIdx]
	delete(this.numToIdx, val)
	
	return true
}

func (this *RandomizedSetWithStats) GetRandom() int {
	this.getRandomCount++
	
	if len(this.nums) == 0 {
		return -1
	}
	
	randomIdx := this.randGen.Intn(len(this.nums))
	return this.nums[randomIdx]
}

func (this *RandomizedSetWithStats) GetStats() (int, int, int) {
	return this.insertCount, this.removeCount, this.getRandomCount
}

func (this *RandomizedSetWithStats) Size() int {
	return len(this.nums)
}

func main() {
	// Test cases
	fmt.Println("=== Testing RandomizedSet ===")
	
	// Test 1: Basic operations
	rs := ConstructorRandomizedSet()
	
	fmt.Printf("Insert 1: %t\n", rs.Insert(1))
	fmt.Printf("Insert 2: %t\n", rs.Insert(2))
	fmt.Printf("Insert 1 again: %t\n", rs.Insert(1))
	fmt.Printf("Current set: %v\n", rs.nums)
	
	fmt.Printf("Remove 1: %t\n", rs.Remove(1))
	fmt.Printf("Remove 1 again: %t\n", rs.Remove(1))
	fmt.Printf("Current set: %v\n", rs.nums)
	
	// Test 2: GetRandom operations
	fmt.Println("\n=== Testing GetRandom ===")
	rs2 := ConstructorRandomizedSet()
	
	for i := 1; i <= 5; i++ {
		rs2.Insert(i)
	}
	
	fmt.Printf("Set: %v\n", rs2.nums)
	
	// Generate some random elements
	fmt.Println("Random elements:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  Random: %d\n", rs2.GetRandom())
	}
	
	// Test 3: Edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	emptyRs := ConstructorRandomizedSet()
	
	fmt.Printf("Insert into empty: %t\n", emptyRs.Insert(10))
	fmt.Printf("Remove from empty: %t\n", emptyRs.Remove(5))
	fmt.Printf("GetRandom from empty: %d\n", emptyRs.GetRandom())
	
	// Test 4: Optimized version
	fmt.Println("\n=== Testing Optimized Version ===")
	optRs := ConstructorRandomizedSetOptimized()
	
	for i := 1; i <= 3; i++ {
		optRs.Insert(i * 10)
	}
	
	fmt.Printf("Set: %v\n", optRs.nums)
	fmt.Printf("Random: %d\n", optRs.GetRandom())
	fmt.Printf("Remove 20: %t\n", optRs.Remove(20))
	fmt.Printf("Set after removal: %v\n", optRs.nums)
	
	// Test 5: Statistics version
	fmt.Println("\n=== Testing Statistics Version ===")
	statsRs := ConstructorRandomizedSetWithStats()
	
	for i := 1; i <= 5; i++ {
		statsRs.Insert(i)
	}
	
	statsRs.Remove(2)
	statsRs.Remove(4)
	
	// Generate some random elements
	for i := 0; i < 5; i++ {
		statsRs.GetRandom()
	}
	
	inserts, removes, randoms := statsRs.GetStats()
	fmt.Printf("Stats - Inserts: %d, Removes: %d, GetRandom: %d\n", inserts, removes, randoms)
	fmt.Printf("Current size: %d\n", statsRs.Size())
	
	// Test 6: Stress test
	fmt.Println("\n=== Stress Test ===")
	stressRs := ConstructorRandomizedSet()
	
	// Insert many elements
	for i := 0; i < 1000; i++ {
		stressRs.Insert(i)
	}
	
	// Remove some elements
	for i := 0; i < 500; i++ {
		stressRs.Remove(i * 2)
	}
	
	// Test random access
	fmt.Printf("Random element from large set: %d\n", stressRs.GetRandom())
	fmt.Printf("Final size: %d\n", len(stressRs.nums))
}
