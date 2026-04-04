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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Array + HashMap for O(1) Randomized Set
- **Array Storage**: Store elements in array for O(1) random access
- **HashMap Indexing**: Track element positions for O(1) insert/delete
- **Swap-Remove**: Remove by swapping with last element
- **Random Selection**: Direct array indexing for uniform distribution

## 2. PROBLEM CHARACTERISTICS
- **O(1) Operations**: Insert, delete, getRandom in constant time
- **Random Access**: Need uniform random element selection
- **Duplicate Prevention**: Set semantics (no duplicates allowed)
- **Dynamic Size**: Support for arbitrary insert/delete operations

## 3. SIMILAR PROBLEMS
- Insert Delete GetRandom O(1) Duplicates (LeetCode 381) - Allow duplicates
- RandomizedCollection (LeetCode 381) - Collection with duplicates
- Design Skip List (LeetCode 1206) - Probabilistic data structure
- Design HashMap (LeetCode 706) - Custom hash map implementation

## 4. KEY OBSERVATIONS
- **Array + HashMap**: Perfect combination for O(1) operations
- **Swap-Remove Technique**: Remove in O(1) by swapping with last element
- **Random Index**: Direct array indexing gives uniform random selection
- **Index Updates**: Must update hashmap when swapping elements

## 5. VARIATIONS & EXTENSIONS
- **Duplicate Support**: Allow duplicate elements in collection
- **Range Queries**: Support for random element in value range
- **Statistics Tracking**: Track insert/delete/getRandom counts
- **Thread Safety**: Add concurrency support with locks

## 6. INTERVIEW INSIGHTS
- Always clarify: "Duplicates allowed? Thread safety? Memory constraints?"
- Edge cases: empty set, single element, large values
- Time complexity: O(1) for all operations
- Space complexity: O(N) where N=number of elements
- Key insight: array + hashmap enables O(1) all operations

## 7. COMMON MISTAKES
- Not updating hashmap indices when swapping elements
- O(N) random access instead of O(1) array indexing
- Not handling empty set edge cases
- Wrong swap-remove logic causing gaps in array
- Not seeding random number generator properly

## 8. OPTIMIZATION STRATEGIES
- **Basic Implementation**: O(1) time, O(N) space - standard
- **Optimized Random**: Better random number generation
- **Statistics Tracking**: O(1) time, O(N) space - with counters
- **Memory Pool**: Reuse memory for multiple instances

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a deck of cards with a quick reference index:**
- You have cards in a specific order (array)
- You have a quick reference telling you where each card is (hashmap)
- When you need a random card, just pick a random position
- When you remove a card, swap it with the last card to keep deck compact
- Reference index helps you find any card instantly

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Set with insert, delete, getRandom operations
2. **Goal**: All operations in O(1) time
3. **Constraints**: No duplicates, uniform random distribution
4. **Output**: Efficient data structure implementation

#### Phase 2: Key Insight Recognition
- **"Array + hashmap natural fit"** → Array for random access, hashmap for O(1) lookup
- **"Swap-remove technique"** → Remove in O(1) by swapping with last element
- **"Random indexing"** → Direct array indexing gives uniform distribution
- **"Index synchronization"** → Must keep hashmap and array in sync

#### Phase 3: Strategy Development
```
Human thought process:
"I need O(1) insert, delete, getRandom.
Regular array/list has O(N) operations.

Array + HashMap Approach:
1. Array: store elements for O(1) random access
2. HashMap: element → index mapping for O(1) lookup
3. Insert: add to array end, update hashmap
4. Delete: find index via hashmap, swap with last, update hashmap
5. GetRandom: generate random index, return array element

This gives O(1) for all operations!"
```

#### Phase 4: Edge Case Handling
- **Empty set**: Return appropriate values (often -1 or error)
- **Single element**: Array and hashmap both have one element
- **Non-existent delete**: Return false for element not found
- **Random from empty**: Handle gracefully

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: Insert(1), Insert(2), Insert(3), Remove(2), GetRandom()

Human thinking:
"Array + HashMap Approach:
1. Insert(1): array=[1], map={1:0}
2. Insert(2): array=[1,2], map={1:0, 2:1}
3. Insert(3): array=[1,2,3], map={1:0, 2:1, 3:2}
4. Remove(2): 
   - Find index: map[2] = 1
   - Swap with last: array[1] = 3, array[2] = 2
   - Update map: map[3] = 1
   - Remove last: array=[1,3], map={1:0, 3:1}
5. GetRandom(): random index 0-1, return array[0] or array[1]

All operations O(1) ✓"
```

#### Phase 6: Intuition Validation
- **Why array works**: Direct indexing gives O(1) random access
- **Why hashmap works**: O(1) lookup for element positions
- **Why swap-remove works**: Maintains array compactness in O(1)
- **Why O(1) all operations**: Each operation touches constant data

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use array?"** → Delete would be O(N) search
2. **"Should I use linked list?"** → Random access would be O(N)
3. **"What about duplicates?"** → Different problem, needs counting
4. **"Can I optimize further?"** → O(1) is already optimal
5. **"What about thread safety?"** → Add locks for concurrent access

### Real-World Analogy
**Like a music playlist with quick track lookup:**
- You have songs in a specific order (array)
- You have a quick reference telling you where each song is (hashmap)
- When you want a random song, just pick a random position
- When you remove a song, swap it with the last song to keep playlist compact
- Reference index helps you find any song instantly
- Like a jukebox with instant song lookup

### Human-Readable Pseudocode
```
class RandomizedSet:
    array = []
    elementToIndex = hashmap()
    
    function insert(val):
        if val in elementToIndex:
            return false
        elementToIndex[val] = len(array)
        array.append(val)
        return true
    
    function remove(val):
        if val not in elementToIndex:
            return false
        idx = elementToIndex[val]
        lastIdx = len(array) - 1
        lastVal = array[lastIdx]
        
        // Swap with last element
        array[idx] = lastVal
        elementToIndex[lastVal] = idx
        
        // Remove last element
        array.pop()
        delete elementToIndex[val]
        return true
    
    function getRandom():
        if len(array) == 0:
            return error
        randomIdx = random(0, len(array) - 1)
        return array[randomIdx]
```

### Execution Visualization

### Example: Operations: Insert(1), Insert(2), Insert(3), Remove(2), GetRandom()
```
Array + HashMap Approach:
Initial: array=[], map={}

Insert(1): array=[1], map={1:0}
Insert(2): array=[1,2], map={1:0, 2:1}
Insert(3): array=[1,2,3], map={1:0, 2:1, 3:2}

Remove(2):
- Find index: map[2] = 1
- Last element: array[2] = 3
- Swap: array[1] = 3, map[3] = 1
- Remove last: array=[1,3], map={1:0, 3:1}

GetRandom():
- Random index 0: return array[0] = 1
- Random index 1: return array[1] = 3

Final state: array=[1,3], map={1:0, 3:1}
All operations O(1) ✓
```

### Key Visualization Points:
- **Array Storage**: Elements stored for O(1) random access
- **HashMap Indexing**: Element → position mapping for O(1) lookup
- **Swap-Remove**: Remove by swapping with last element
- **Random Selection**: Direct array indexing for uniform distribution

### Memory Layout Visualization:
```
State Evolution:
Initial: array=[], map={}

After Insert(1): array=[1], map={1:0}
After Insert(2): array=[1,2], map={1:0, 2:1}
After Insert(3): array=[1,2,3], map={1:0, 2:1, 3:2}

Remove(2) Process:
1. Find: map[2] = 1
2. Last: array[2] = 3
3. Swap: array[1] = 3, map[3] = 1
4. Remove: array=[1,3], map={1:0, 3:1}

Operation Complexity:
Insert: O(1) time, O(1) space
Remove: O(1) time, O(1) space
GetRandom: O(1) time, O(1) space
```

### Time Complexity Breakdown:
- **Insert**: O(1) time (hashmap lookup + array append), O(1) space
- **Remove**: O(1) time (hashmap lookup + swap + pop), O(1) space
- **GetRandom**: O(1) time (random number + array access), O(1) space
- **Space**: O(N) where N=number of elements

### Alternative Approaches:

#### 1. Linked List + HashMap (O(1) average, O(N) worst)
```go
type RandomizedSetLinkedList struct {
    head *Node
    elementToNode map[int]*Node
}
```
- **Pros**: Memory efficient for insertions
- **Cons**: GetRandom requires traversal, O(N) time

#### 2. Balanced BST (O(log N) time, O(N) space)
```go
type RandomizedSetBST struct {
    tree *AVLTree
    size int
}
```
- **Pros**: Ordered elements, range queries
- **Cons**: GetRandom requires O(N) traversal or extra storage

#### 3. Skip List (O(log N) time, O(N) space)
```go
type RandomizedSetSkipList struct {
    head *SkipListNode
    size int
}
```
- **Pros**: Good for range queries, probabilistic balance
- **Cons**: More complex implementation

### Extensions for Interviews:
- **Duplicate Support**: Allow multiple instances of same value
- **Thread Safety**: Add mutex locks for concurrent access
- **Statistics**: Track operation counts and patterns
- **Range Queries**: Support for random element in value range
- **Memory Optimization**: Use object pooling for frequent operations
*/
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
