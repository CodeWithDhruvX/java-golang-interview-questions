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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Array + HashMap with Duplicate Support
- **Array Storage**: Store elements for O(1) random access
- **HashMap with Lists**: Value → list of indices mapping for duplicates
- **Swap-Remove**: Remove by swapping with last element
- **Index Management**: Update all indices when elements move

## 2. PROBLEM CHARACTERISTICS
- **O(1) Operations**: Insert, delete, getRandom in constant time
- **Duplicate Support**: Allow multiple instances of same value
- **Random Access**: Need uniform random element selection
- **Dynamic Size**: Support for arbitrary insert/delete operations

## 3. SIMILAR PROBLEMS
- Insert Delete GetRandom O(1) (LeetCode 380) - No duplicates version
- RandomizedCollection (LeetCode 381) - Same problem name
- Design Skip List (LeetCode 1206) - Probabilistic data structure
- Design HashMap (LeetCode 706) - Custom hash map implementation

## 4. KEY OBSERVATIONS
- **Multiple Indices**: One value can have multiple array indices
- **Index Updates**: Must update all indices when elements move
- **Swap Complexity**: More complex due to multiple index updates
- **Memory Trade-off**: More memory for duplicate support

## 5. VARIATIONS & EXTENSIONS
- **Count Tracking**: Track count of each value instead of indices
- **Optimized Indexing**: Use bitsets or other optimized structures
- **Statistics Tracking**: Track operation counts and patterns
- **Thread Safety**: Add concurrency support with locks

## 6. INTERVIEW INSIGHTS
- Always clarify: "Duplicate count limit? Remove one or all? Thread safety?"
- Edge cases: empty collection, single element, many duplicates
- Time complexity: O(1) for all operations
- Space complexity: O(N) where N=number of elements
- Key insight: value→indices mapping enables duplicate handling

## 7. COMMON MISTAKES
- Not updating all indices when swapping elements
- Wrong duplicate removal logic (remove one vs all)
- Index synchronization errors between array and hashmap
- Memory leaks in index list management
- Not handling empty collection edge cases

## 8. OPTIMIZATION STRATEGIES
- **Basic Implementation**: O(1) time, O(N) space - standard
- **Optimized Indexing**: Use bitsets for index tracking
- **Count-Based**: O(1) time, O(N) space - with value counts
- **Statistics Tracking**: O(1) time, O(N) space - with counters

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a deck of cards with multiple copies and a detailed index:**
- You have cards where some cards have multiple copies (array)
- You have a detailed reference telling you all positions of each card type (hashmap with lists)
- When you need a random card, just pick a random position
- When you remove a card, swap it with the last card and update all references
- Reference index helps you find any card type and all its positions

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Collection with insert, delete, getRandom operations
2. **Goal**: All operations in O(1) time, duplicates allowed
3. **Constraints**: Multiple instances of same value allowed
4. **Output**: Efficient data structure with duplicate support

#### Phase 2: Key Insight Recognition
- **"Value→indices mapping"** → One value maps to multiple array positions
- **"Index synchronization"** → Must update all indices when elements move
- **"Swap-remove complexity"** → More complex due to multiple index updates
- **"Random access unchanged"** → Still O(1) with array indexing

#### Phase 3: Strategy Development
```
Human thought process:
"I need O(1) insert, delete, getRandom with duplicates.
Regular array+hashmap won't handle duplicates properly.

Array + HashMap with Lists Approach:
1. Array: store elements for O(1) random access
2. HashMap: value → list of indices mapping
3. Insert: add to array end, append index to value's list
4. Delete: find any index of value, swap with last, update all affected indices
5. GetRandom: generate random index, return array element

This handles duplicates correctly!"
```

#### Phase 4: Edge Case Handling
- **Empty collection**: Return appropriate values (often -1 or error)
- **Single element**: Array and hashmap both have one element
- **All duplicates**: All elements have same value
- **Non-existent delete**: Return false for value not found

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: Insert(1), Insert(1), Insert(2), Remove(1), GetRandom()

Human thinking:
"Array + HashMap with Lists Approach:
1. Insert(1): array=[1], map={1:[0]}
2. Insert(1): array=[1,1], map={1:[0,1]}
3. Insert(2): array=[1,1,2], map={1:[0,1], 2:[2]}
4. Remove(1): 
   - Find any index: map[1][0] = 0
   - Last element: array[2] = 2
   - Swap: array[0] = 2, map[2] = [0]
   - Remove last: array=[2,1], map={1:[1], 2:[0]}
5. GetRandom(): random index 0-1, return array[0] or array[1]

All operations O(1) ✓"
```

#### Phase 6: Intuition Validation
- **Why index lists work**: Track all positions of each value
- **Why swap-remove works**: Maintains array compactness in O(1)
- **Why index updates needed**: Moving one element affects other value indices
- **Why O(1) all operations**: Each operation touches constant data

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use array?"** → Delete would be O(N) search for duplicates
2. **"Should I use multiset?"** → Different semantics, need specific operations
3. **"What about removing all duplicates?"** → Clarify removal semantics
4. **"Can I optimize further?"** → O(1) is already optimal
5. **"What about thread safety?"** → Add locks for concurrent access

### Real-World Analogy
**Like a music library with multiple copies of songs:**
- You have songs where some songs have multiple copies (array)
- You have a detailed index telling you all positions of each song (hashmap with lists)
- When you want a random song, just pick a random position
- When you remove a song copy, swap it with the last song and update all references
- Reference index helps you find any song type and all its positions
- Like a streaming service with multiple copies of popular songs

### Human-Readable Pseudocode
```
class RandomizedCollection:
    array = []
    valueToIndices = hashmap() // value -> list of indices
    
    function insert(val):
        array.append(val)
        if val not in valueToIndices:
            valueToIndices[val] = []
        valueToIndices[val].append(len(array) - 1)
        return true
    
    function remove(val):
        if val not in valueToIndices or valueToIndices[val].isEmpty():
            return false
        
        // Get any index of val to remove
        idxToRemove = valueToIndices[val].pop()
        lastIdx = len(array) - 1
        lastVal = array[lastIdx]
        
        // Swap with last element
        array[idxToRemove] = lastVal
        
        // Update indices of lastVal
        for i, idx in enumerate(valueToIndices[lastVal]):
            if idx == lastIdx:
                valueToIndices[lastVal][i] = idxToRemove
                break
        
        // Remove last element
        array.pop()
        return true
    
    function getRandom():
        if len(array) == 0:
            return error
        randomIdx = random(0, len(array) - 1)
        return array[randomIdx]
```

### Execution Visualization

### Example: Operations: Insert(1), Insert(1), Insert(2), Remove(1), GetRandom()
```
Array + HashMap with Lists Approach:
Initial: array=[], map={}

Insert(1): array=[1], map={1:[0]}
Insert(1): array=[1,1], map={1:[0,1]}
Insert(2): array=[1,1,2], map={1:[0,1], 2:[2]}

Remove(1):
- Find any index: map[1].pop() = 0
- Last element: array[2] = 2
- Swap: array[0] = 2, map[2] = [0]
- Remove last: array=[2,1], map={1:[1], 2:[0]}

GetRandom():
- Random index 0: return array[0] = 2
- Random index 1: return array[1] = 1

Final state: array=[2,1], map={1:[1], 2:[0]}
All operations O(1) ✓
```

### Key Visualization Points:
- **Array Storage**: Elements stored for O(1) random access
- **Index Lists**: Value → list of indices mapping for duplicates
- **Swap-Remove**: Remove by swapping with last element
- **Index Updates**: Must update all affected indices when swapping

### Memory Layout Visualization:
```
State Evolution:
Initial: array=[], map={}

After Insert(1): array=[1], map={1:[0]}
After Insert(1): array=[1,1], map={1:[0,1]}
After Insert(2): array=[1,1,2], map={1:[0,1], 2:[2]}

Remove(1) Process:
1. Find: map[1].pop() = 0
2. Last: array[2] = 2
3. Swap: array[0] = 2, map[2] = [0]
4. Remove: array=[2,1], map={1:[1], 2:[0]}

Operation Complexity:
Insert: O(1) time, O(1) space
Remove: O(1) time, O(1) space (amortized)
GetRandom: O(1) time, O(1) space
```

### Time Complexity Breakdown:
- **Insert**: O(1) time (hashmap lookup + array append), O(1) space
- **Remove**: O(1) time (hashmap lookup + swap + updates), O(1) space
- **GetRandom**: O(1) time (random number + array access), O(1) space
- **Space**: O(N) where N=number of elements

### Alternative Approaches:

#### 1. Count-Based Approach (O(1) time, O(N) space)
```go
type RandomizedCollectionCount struct {
    valueToCount map[int]int
    values       []int
}
```
- **Pros**: Simpler index management
- **Cons**: Cannot support specific element removal easily

#### 2. Two-Level HashMap (O(1) time, O(N) space)
```go
type RandomizedCollectionTwoLevel struct {
    array         []int
    valueToIndices map[int]map[int]bool
}
```
- **Pros**: Faster index lookup for duplicates
- **Cons**: More complex implementation

#### 3. Segmented Array (O(1) time, O(N) space)
```go
type RandomizedCollectionSegmented struct {
    segments [][]int
    valueToSegment map[int][]int
}
```
- **Pros**: Better cache locality
- **Cons**: More complex random access

### Extensions for Interviews:
- **Remove All**: Support for removing all instances of a value
- **Count Queries**: Support for getting count of specific value
- **Range Queries**: Support for random element in value range
- **Thread Safety**: Add mutex locks for concurrent access
- **Memory Optimization**: Use object pooling for frequent operations
*/
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
