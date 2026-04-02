package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 380. Insert Delete GetRandom O(1) - Duplicates allowed
// Time: O(1) for all operations, Space: O(N)
type RandomizedCollection struct {
	nums     []int
	valToIdx map[int][]int // value -> list of indices
}

// Constructor initializes the data structure
func ConstructorRandomizedCollection() RandomizedCollection {
	return RandomizedCollection{
		nums:     []int{},
		valToIdx: make(map[int][]int),
	}
}

// Insert inserts a value to the collection. Returns true if the collection did not contain the specified element.
func (this *RandomizedCollection) Insert(val int) bool {
	hasValue := len(this.valToIdx[val]) > 0
	
	// Add value to array
	this.nums = append(this.nums, val)
	
	// Add index to map
	this.valToIdx[val] = append(this.valToIdx[val], len(this.nums)-1)
	
	return !hasValue
}

// Remove removes a value from the collection. Returns true if the collection contained the specified element.
func (this *RandomizedCollection) Remove(val int) bool {
	if len(this.valToIdx[val]) == 0 {
		return false
	}
	
	// Get index of element to remove
	idxToRemove := this.valToIdx[val][len(this.valToIdx[val])-1]
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	// Move last element to the position of element to remove
	this.nums[idxToRemove] = lastVal
	
	// Update indices in map
	// Remove old index from val's list
	this.valToIdx[val] = this.valToIdx[val][:len(this.valToIdx[val])-1]
	
	// Update last value's indices
	if len(this.valToIdx[lastVal]) > 0 {
		// Find and update the index of lastVal
		for i, idx := range this.valToIdx[lastVal] {
			if idx == lastIdx {
				this.valToIdx[lastVal][i] = idxToRemove
				break
			}
		}
	}
	
	// Remove last element from array
	this.nums = this.nums[:lastIdx]
	
	// Clean up empty lists
	if len(this.valToIdx[val]) == 0 {
		delete(this.valToIdx, val)
	}
	
	return true
}

// GetRandom returns a random element from the current collection.
func (this *RandomizedCollection) GetRandom() int {
	if len(this.nums) == 0 {
		return -1 // Or handle error appropriately
	}
	
	rand.Seed(time.Now().UnixNano())
	randomIdx := rand.Intn(len(this.nums))
	return this.nums[randomIdx]
}

// Alternative implementation with better index management
type RandomizedCollectionOptimized struct {
	nums     []int
	valToIdx map[int]map[int]bool // value -> set of indices
	randGen  *rand.Rand
}

func ConstructorRandomizedCollectionOptimized() RandomizedCollectionOptimized {
	source := rand.NewSource(time.Now().UnixNano())
	return RandomizedCollectionOptimized{
		nums:     []int{},
		valToIdx: make(map[int]map[int]bool),
		randGen:  rand.New(source),
	}
}

func (this *RandomizedCollectionOptimized) Insert(val int) bool {
	hasValue := len(this.valToIdx[val]) > 0
	
	this.nums = append(this.nums, val)
	
	if this.valToIdx[val] == nil {
		this.valToIdx[val] = make(map[int]bool)
	}
	this.valToIdx[val][len(this.nums)-1] = true
	
	return !hasValue
}

func (this *RandomizedCollectionOptimized) Remove(val int) bool {
	if len(this.valToIdx[val]) == 0 {
		return false
	}
	
	// Get any index of val to remove
	var idxToRemove int
	for idx := range this.valToIdx[val] {
		idxToRemove = idx
		break
	}
	
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	// Move last element to the position of element to remove
	this.nums[idxToRemove] = lastVal
	
	// Update indices
	delete(this.valToIdx[val], idxToRemove)
	
	// Update last value's indices
	if len(this.valToIdx[lastVal]) > 0 {
		delete(this.valToIdx[lastVal], lastIdx)
		this.valToIdx[lastVal][idxToRemove] = true
	}
	
	// Remove last element
	this.nums = this.nums[:lastIdx]
	
	// Clean up empty maps
	if len(this.valToIdx[val]) == 0 {
		delete(this.valToIdx, val)
	}
	
	return true
}

func (this *RandomizedCollectionOptimized) GetRandom() int {
	if len(this.nums) == 0 {
		return -1
	}
	
	randomIdx := this.randGen.Intn(len(this.nums))
	return this.nums[randomIdx]
}

// Version with detailed tracking and statistics
type RandomizedCollectionWithStats struct {
	nums          []int
	valToIdx      map[int][]int
	insertCount   int
	removeCount   int
	getRandomCount int
	duplicateCount int
	randGen       *rand.Rand
}

func ConstructorRandomizedCollectionWithStats() RandomizedCollectionWithStats {
	source := rand.NewSource(time.Now().UnixNano())
	return RandomizedCollectionWithStats{
		nums:     []int{},
		valToIdx: make(map[int][]int),
		randGen:  rand.New(source),
	}
}

func (this *RandomizedCollectionWithStats) Insert(val int) bool {
	this.insertCount++
	
	hasValue := len(this.valToIdx[val]) > 0
	
	this.nums = append(this.nums, val)
	this.valToIdx[val] = append(this.valToIdx[val], len(this.nums)-1)
	
	if hasValue {
		this.duplicateCount++
	}
	
	return !hasValue
}

func (this *RandomizedCollectionWithStats) Remove(val int) bool {
	this.removeCount++
	
	if len(this.valToIdx[val]) == 0 {
		return false
	}
	
	idxToRemove := this.valToIdx[val][len(this.valToIdx[val])-1]
	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]
	
	this.nums[idxToRemove] = lastVal
	
	this.valToIdx[val] = this.valToIdx[val][:len(this.valToIdx[val])-1]
	
	if len(this.valToIdx[lastVal]) > 0 {
		for i, idx := range this.valToIdx[lastVal] {
			if idx == lastIdx {
				this.valToIdx[lastVal][i] = idxToRemove
				break
			}
		}
	}
	
	this.nums = this.nums[:lastIdx]
	
	if len(this.valToIdx[val]) == 0 {
		delete(this.valToIdx, val)
	}
	
	return true
}

func (this *RandomizedCollectionWithStats) GetRandom() int {
	this.getRandomCount++
	
	if len(this.nums) == 0 {
		return -1
	}
	
	randomIdx := this.randGen.Intn(len(this.nums))
	return this.nums[randomIdx]
}

func (this *RandomizedCollectionWithStats) GetStats() (int, int, int, int) {
	return this.insertCount, this.removeCount, this.getRandomCount, this.duplicateCount
}

func (this *RandomizedCollectionWithStats) Size() int {
	return len(this.nums)
}

func (this *RandomizedCollectionWithStats) UniqueCount() int {
	return len(this.valToIdx)
}

func main() {
	// Test cases
	fmt.Println("=== Testing RandomizedCollection ===")
	
	// Test 1: Basic operations with duplicates
	rc := ConstructorRandomizedCollection()
	
	fmt.Printf("Insert 1: %t\n", rc.Insert(1))
	fmt.Printf("Insert 1 again: %t\n", rc.Insert(1))
	fmt.Printf("Insert 2: %t\n", rc.Insert(2))
	fmt.Printf("Current collection: %v\n", rc.nums)
	
	fmt.Printf("Remove 1: %t\n", rc.Remove(1))
	fmt.Printf("Current collection: %v\n", rc.nums)
	
	// Test 2: GetRandom with duplicates
	fmt.Println("\n=== Testing GetRandom with Duplicates ===")
	rc2 := ConstructorRandomizedCollection()
	
	// Insert duplicates
	for i := 0; i < 3; i++ {
		rc2.Insert(10)
	}
	rc2.Insert(20)
	
	fmt.Printf("Collection: %v\n", rc2.nums)
	fmt.Println("Random elements:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  Random: %d\n", rc2.GetRandom())
	}
	
	// Test 3: Edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	emptyRc := ConstructorRandomizedCollection()
	
	fmt.Printf("Insert into empty: %t\n", emptyRc.Insert(5))
	fmt.Printf("Remove from empty: %t\n", emptyRc.Remove(3))
	fmt.Printf("GetRandom from empty: %d\n", emptyRc.GetRandom())
	
	// Test 4: Optimized version
	fmt.Println("\n=== Testing Optimized Version ===")
	optRc := ConstructorRandomizedCollectionOptimized()
	
	optRc.Insert(100)
	optRc.Insert(100)
	optRc.Insert(200)
	optRc.Insert(100)
	
	fmt.Printf("Collection: %v\n", optRc.nums)
	fmt.Printf("Random: %d\n", optRc.GetRandom())
	fmt.Printf("Remove 100: %t\n", optRc.Remove(100))
	fmt.Printf("Collection after removal: %v\n", optRc.nums)
	
	// Test 5: Statistics version
	fmt.Println("\n=== Testing Statistics Version ===")
	statsRc := ConstructorRandomizedCollectionWithStats()
	
	// Insert many elements with duplicates
	for i := 1; i <= 5; i++ {
		statsRc.Insert(i)
		statsRc.Insert(i) // Duplicate
	}
	
	statsRc.Remove(2)
	statsRc.Remove(4)
	
	// Generate random elements
	for i := 0; i < 5; i++ {
		statsRc.GetRandom()
	}
	
	inserts, removes, randoms, duplicates := statsRc.GetStats()
	fmt.Printf("Stats - Inserts: %d, Removes: %d, GetRandom: %d, Duplicates: %d\n", 
		inserts, removes, randoms, duplicates)
	fmt.Printf("Current size: %d, Unique count: %d\n", statsRc.Size(), statsRc.UniqueCount())
	
	// Test 6: Stress test with many duplicates
	fmt.Println("\n=== Stress Test with Duplicates ===")
	stressRc := ConstructorRandomizedCollection()
	
	// Insert many elements with duplicates
	for i := 0; i < 100; i++ {
		stressRc.Insert(i % 10) // Only 0-9, but 100 elements total
	}
	
	fmt.Printf("Stress test - Size: %d, Unique: %d\n", stressRc.Size(), len(stressRc.valToIdx))
	fmt.Printf("Random element: %d\n", stressRc.GetRandom())
	
	// Remove some elements
	for i := 0; i < 50; i++ {
		stressRc.Remove(i % 5)
	}
	
	fmt.Printf("After removals - Size: %d, Unique: %d\n", stressRc.Size(), len(stressRc.valToIdx))
}
