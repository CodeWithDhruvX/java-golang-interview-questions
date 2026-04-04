package main

import "fmt"

// 128. Longest Consecutive Sequence
// Time: O(N), Space: O(N)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	
	maxLength := 0
	
	for num := range numSet {
		// Only start counting from the beginning of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1
			
			// Count the length of the consecutive sequence
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}
			
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	
	return maxLength
}

// Union-Find approach
func longestConsecutiveUnionFind(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Initialize Union-Find
	parent := make(map[int]int)
	rank := make(map[int]int)
	
	for _, num := range nums {
		parent[num] = num
		rank[num] = 0
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	var union func(int, int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return
		}
		
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
	}
	
	// Union consecutive numbers
	for _, num := range nums {
		if _, exists := parent[num-1]; exists {
			union(num, num-1)
		}
		if _, exists := parent[num+1]; exists {
			union(num, num+1)
		}
	}
	
	// Count the size of each component
	componentSize := make(map[int]int)
	for _, num := range nums {
		root := find(num)
		componentSize[root]++
	}
	
	maxLength := 0
	for _, size := range componentSize {
		if size > maxLength {
			maxLength = size
		}
	}
	
	return maxLength
}

// Sorting approach (O(N log N))
func longestConsecutiveSorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Remove duplicates and sort
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}
	
	sorted := make([]int, 0, len(unique))
	for num := range unique {
		sorted = append(sorted, num)
	}
	
	// Simple bubble sort for demonstration (in practice, use sort.Ints)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	maxLength := 1
	currentLength := 1
	
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1]+1 {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}
	
	if currentLength > maxLength {
		maxLength = currentLength
	}
	
	return maxLength
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: HashSet-based Sequence Detection
- **HashSet Storage**: Store all numbers for O(1) lookup
- **Sequence Start Detection**: Only start counting from sequence beginnings
- **Linear Expansion**: Expand forward from each sequence start
- **Maximum Tracking**: Track longest sequence found

## 2. PROBLEM CHARACTERISTICS
- **Unordered Array**: Input array with no specific order
- **Consecutive Numbers**: Find longest sequence of consecutive integers
- **Unique Elements**: Duplicates don't affect sequence length
- **Linear Time**: Need O(N) solution (sorting would be O(N log N))

## 3. SIMILAR PROBLEMS
- Longest Consecutive Sequence in Binary Tree (LeetCode 298)
- Consecutive Numbers Sum (LeetCode - variant)
- Find Missing Numbers (LeetCode - variant)
- Array Partition into Consecutive Subsets (LeetCode - variant)

## 4. KEY OBSERVATIONS
- **HashSet advantage**: O(1) lookup for number existence
- **Sequence start condition**: num-1 not in set indicates sequence start
- **Forward expansion**: Only need to check num+1, num+2, etc.
- **Duplicate handling**: HashSet automatically handles duplicates

## 5. VARIATIONS & EXTENSIONS
- **Longest Consecutive Subarray**: Consecutive in original order
- **K-length Consecutive Sequences**: Find all sequences of length K
- **Circular Consecutive**: Treat array as circular
- **Multiple Arrays**: Find longest across multiple arrays

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can the array be empty? Are there duplicates?"
- Edge cases: empty array, single element, all duplicates
- Time complexity: O(N) - each number processed once
- Space complexity: O(N) - for HashSet storage

## 7. COMMON MISTAKES
- Not using HashSet and trying O(N²) approach
- Starting sequence from every number instead of just starts
- Not handling empty array case
- Forgetting to handle duplicates properly
- Using sorting (O(N log N)) when O(N) is possible

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(N) space
- **Union Find alternative**: More complex but demonstrates pattern
- **Bit manipulation**: Not applicable for this problem
- **Early termination**: Not applicable (need to check all numbers)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest chain of consecutive numbers:**
- You have a bag of numbered balls (unordered array)
- You want to find the longest chain where each number is one more than the previous
- Instead of checking every possible starting point, only start chains at the beginning
- A number is the beginning if the previous number doesn't exist
- Once you find a beginning, count how long the chain goes forward

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Unordered array of integers
2. **Goal**: Find length of longest consecutive sequence
3. **Output**: Integer representing maximum sequence length
4. **Constraint**: Must be O(N) time complexity

#### Phase 2: Key Insight Recognition
- **"HashSet necessity"** → Need O(1) lookup to achieve O(N) time
- **"Sequence start detection"** → Only start from numbers where num-1 doesn't exist
- **"Forward expansion"** → Once at sequence start, count forward consecutively
- **"Duplicate handling"** → HashSet automatically handles duplicates

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the longest consecutive sequence.
If I sort, it's O(N log N), but I need O(N).
I'll use a HashSet for O(1) lookups.
For each number, I'll check if it's the start of a sequence.
If num-1 doesn't exist, it's a sequence start.
From there, I'll count how far the sequence goes forward."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **Single element**: Return 1
- **All duplicates**: Return 1
- **All consecutive**: Return array length (deduplicated)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [100, 4, 200, 1, 3, 2]

Human thinking:
"I'll build a HashSet: {100, 4, 200, 1, 3, 2}

Now check each number:
100: Is 99 in set? No, so 100 is sequence start
- Count: 100 (101? No) → length = 1

4: Is 3 in set? Yes, so 4 is not sequence start

200: Is 199 in set? No, so 200 is sequence start  
- Count: 200 (201? No) → length = 1

1: Is 0 in set? No, so 1 is sequence start
- Count: 1, 2, 3, 4 (5? No) → length = 4

3: Is 2 in set? Yes, so 3 is not sequence start
2: Is 1 in set? Yes, so 2 is not sequence start

Longest sequence: [1,2,3,4] with length 4"
```

#### Phase 6: Intuition Validation
- **Why HashSet works**: Enables O(1) existence checks
- **Why start detection works**: Avoids redundant counting from middle of sequences
- **Why O(N) time**: Each number processed at most twice (once as potential start, once in expansion)
- **Why O(N) space**: HashSet stores all unique numbers

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → Sorting is O(N log N), need O(N) solution
2. **"Should I use Union Find?"** → Can, but HashSet is simpler and more efficient
3. **"What about very large numbers?"** → HashSet handles any integer range
4. **"Can I optimize further?"** → Already optimal O(N) time solution

### Real-World Analogy
**Like finding the longest consecutive streak in attendance records:**
- You have attendance records for various days (unordered dates)
- You want to find the longest streak of consecutive days attended
- Instead of checking every day as a potential start, only start when previous day was absent
- Once you find a streak start, count how many consecutive days follow
- Keep track of the longest streak found

### Human-Readable Pseudocode
```
function longestConsecutive(nums):
    if nums is empty:
        return 0
    
    numSet = set(nums)
    maxLength = 0
    
    for num in numSet:
        if num-1 not in numSet:  // Only start from sequence beginnings
            currentNum = num
            currentLength = 1
            
            while currentNum+1 in numSet:
                currentNum++
                currentLength++
            
            maxLength = max(maxLength, currentLength)
    
    return maxLength
```

### Execution Visualization

### Example Array: [100, 4, 200, 1, 3, 2]
```
HashSet: {1, 2, 3, 4, 100, 200}

Processing:
100: 99 not in set → start sequence [100] → length 1
4:   3 in set → skip
200: 199 not in set → start sequence [200] → length 1  
1:   0 not in set → start sequence [1,2,3,4] → length 4
3:   2 in set → skip
2:   1 in set → skip

Maximum length: 4
```

### Key Visualization Points:
- **HashSet creation**: All numbers stored for O(1) lookup
- **Sequence start detection**: Only process numbers where num-1 doesn't exist
- **Forward expansion**: Count consecutive numbers from each start
- **Maximum tracking**: Keep longest sequence found

### Memory Layout Visualization:
```
HashSet: {1, 2, 3, 4, 100, 200}
             ↓     ↓     ↓
         starts: 1, 100, 200
         sequences: [1,2,3,4], [100], [200]
         max_length: 4
```

### Time Complexity Breakdown:
- **HashSet creation**: O(N) time, O(N) space
- **Sequence detection**: Each number processed at most twice
- **Total time**: O(N) where N is number of elements
- **Space**: O(N) for HashSet storage

### Alternative Approaches:

#### 1. Union Find Approach (O(N α(N)) time, O(N) space)
```go
func longestConsecutiveUnionFind(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Initialize Union-Find
    parent := make(map[int]int)
    rank := make(map[int]int)
    
    for _, num := range nums {
        parent[num] = num
        rank[num] = 0
    }
    
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    
    var union func(int, int)
    union = func(x, y int) {
        rootX := find(x)
        rootY := find(y)
        
        if rootX == rootY {
            return
        }
        
        if rank[rootX] < rank[rootY] {
            parent[rootX] = rootY
        } else if rank[rootX] > rank[rootY] {
            parent[rootY] = rootX
        } else {
            parent[rootY] = rootX
            rank[rootX]++
        }
    }
    
    // Union consecutive numbers
    for _, num := range nums {
        if _, exists := parent[num-1]; exists {
            union(num, num-1)
        }
        if _, exists := parent[num+1]; exists {
            union(num, num+1)
        }
    }
    
    // Count the size of each component
    componentSize := make(map[int]int)
    for _, num := range nums {
        root := find(num)
        componentSize[root]++
    }
    
    maxLength := 0
    for _, size := range componentSize {
        if size > maxLength {
            maxLength = size
        }
    }
    
    return maxLength
}
```
- **Pros**: Demonstrates Union Find pattern, good for dynamic queries
- **Cons**: More complex, higher constant factors

#### 2. Sorting Approach (O(N log N) time, O(1) space)
```go
func longestConsecutiveSorting(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Remove duplicates and sort
    unique := make(map[int]bool)
    for _, num := range nums {
        unique[num] = true
    }
    
    sorted := make([]int, 0, len(unique))
    for num := range unique {
        sorted = append(sorted, num)
    }
    
    // Sort the array
    sort.Ints(sorted)
    
    maxLength := 1
    currentLength := 1
    
    for i := 1; i < len(sorted); i++ {
        if sorted[i] == sorted[i-1]+1 {
            currentLength++
        } else {
            if currentLength > maxLength {
                maxLength = currentLength
            }
            currentLength = 1
        }
    }
    
    if currentLength > maxLength {
        maxLength = currentLength
    }
    
    return maxLength
}
```
- **Pros**: Simple to understand and implement
- **Cons**: O(N log N) time, not optimal for large inputs

#### 3. Bitset Approach (O(N) time, O(R) space where R is range)
```go
func longestConsecutiveBitset(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Find min and max to determine range
    minNum, maxNum := nums[0], nums[0]
    for _, num := range nums {
        if num < minNum {
            minNum = num
        }
        if num > maxNum {
            maxNum = num
        }
    }
    
    // Create bitset (simplified - in practice use proper bitset)
    present := make([]bool, maxNum-minNum+1)
    for _, num := range nums {
        present[num-minNum] = true
    }
    
    maxLength := 0
    currentLength := 0
    
    for i := 0; i < len(present); i++ {
        if present[i] {
            currentLength++
            if currentLength > maxLength {
                maxLength = currentLength
            }
        } else {
            currentLength = 0
        }
    }
    
    return maxLength
}
```
- **Pros**: O(N) time, efficient for small number ranges
- **Cons**: O(R) space where R is number range, not suitable for large ranges

### Extensions for Interviews:
- **Longest Consecutive Subarray**: Maintain consecutive in original order
- **K-length Sequences**: Find all sequences of exactly length K
- **Multiple Queries**: Handle multiple longest sequence queries efficiently
- **Circular Array**: Treat array as circular for consecutive sequences
- **Dynamic Updates**: Handle insertions and deletions efficiently
*/
func main() {
	// Test cases
	testCases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		{},
		{1},
		{1, 2, 0, 1},
		{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 3, 5, 7, 9},
		{-1, -2, -3, 0, 1, 2, 3},
	}
	
	for i, nums := range testCases {
		result1 := longestConsecutive(nums)
		result2 := longestConsecutiveUnionFind(nums)
		result3 := longestConsecutiveSorting(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  HashSet: %d, Union-Find: %d, Sorting: %d\n\n", result1, result2, result3)
	}
}
