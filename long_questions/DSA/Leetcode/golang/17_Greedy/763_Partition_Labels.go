package main

import "fmt"

// 763. Partition Labels
// Time: O(N), Space: O(1) for 26 letters
func partitionLabels(s string) []int {
	// Record the last occurrence of each character
	last := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}
	
	var result []int
	start := 0
	end := 0
	
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]])
		
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Greedy with Last Occurrence Tracking
- **Greedy Strategy**: Track last occurrence of each character
- **Partition Points**: Create partitions when current character repeats
- **Single Pass**: Process string once from left to right
- **Optimal Partitioning**: Each partition contains unique characters

## 2. PROBLEM CHARACTERISTICS
- **String Partitioning**: Divide string into maximum number of partitions
- **Unique Character Constraint**: Each partition must contain unique characters
- **Greedy Validity**: Making partitions as early as possible is optimal
- **Character Tracking**: Need to know last occurrence of each character

## 3. SIMILAR PROBLEMS
- Jump Game (LeetCode 55) - Greedy with range tracking
- Gas Station (LeetCode 134) - Circular route feasibility
- Candy (LeetCode 135) - Distribution with constraints
- Merge Intervals (LeetCode 56) - Interval merging

## 4. KEY OBSERVATIONS
- **Last Occurrence**: Each character's last position determines partition boundary
- **Greedy Partition**: Create partition when current character repeats
- **Optimal Strategy**: Making partitions earlier allows more partitions
- **Character Uniqueness**: Each partition must have unique characters

## 5. VARIATIONS & EXTENSIONS
- **Different Character Sets**: Unicode, case sensitivity
- **Partition Constraints**: Maximum/minimum partition sizes
- **Multiple Strings**: Apply same logic to multiple strings
- **Character Weights**: Different importance for characters

## 6. INTERVIEW INSIGHTS
- Always clarify: "What character set? Case sensitivity?"
- Edge cases: empty string, single character, all unique/all same
- Time complexity: O(N) time, O(1) space (26 letters)
- Key insight: track last occurrence to determine partition points
- Greedy works because earlier partitions allow more total partitions

## 7. COMMON MISTAKES
- Not tracking last occurrence correctly
- Wrong partition boundary calculation
- Not handling edge cases (empty string, single character)
- Using complex data structures when simple array suffices
- Off-by-one errors in string indexing

## 8. OPTIMIZATION STRATEGIES
- **Single Pass**: O(N) time, O(1) space - optimal
- **Character Array**: Use fixed-size array for 26 letters
- **Early Partition**: Create partitions as soon as character repeats
- **Last Occurrence**: Track to determine optimal partition boundaries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like dividing a book into chapters with unique character sets:**
- You have a book (string) and want to divide into chapters
- Each chapter must contain unique characters (no repeats)
- You want maximum number of chapters
- When you see a character that appeared before, you must start new chapter
- Track where each character last appeared to know when to split

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s
2. **Goal**: Partition into maximum parts with unique characters
3. **Constraint**: Each part must contain no repeated characters
4. **Output**: Array of partition lengths

#### Phase 2: Key Insight Recognition
- **"Greedy natural fit"** → Make partitions as early as possible
- **"Last occurrence tracking"** → Need to know where characters last appeared
- **"Character uniqueness"** → Each partition must have unique characters
- **"Optimal partitioning"** → Earlier partitions allow more total partitions

#### Phase 3: Strategy Development
```
Human thought process:
"I need to partition string with unique characters in each part.
I'll track the last position where each character appears.
As I scan the string, when I see a character that appeared before,
I need to start a new partition.
The partition boundary is the maximum of all last occurrences seen so far.
This ensures each partition has unique characters."
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty array
- **Single character**: Return [1]
- **All unique**: Each character gets its own partition
- **All same**: Only one partition possible

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
s = "ababcbacadefegdehijhklij"

Human thinking:
"I'll scan and track last occurrence of each character:
Initialize last array with -1 for all 26 letters.

Position 0: 'a'
- last['a'] = -1 (not seen before)
- Update last['a'] = 0
- Current partition: [0,0]

Position 1: 'b'
- last['b'] = -1 (not seen before)
- Update last['b'] = 1
- Current partition: [0,1]

Position 2: 'a'
- last['a'] = 0 (seen before at position 0)
- Need to partition at max(0, last['a']) = 0
- Partition: [0,1] (positions 0-1)
- Update last['a'] = 2
- Start new partition at 2

Position 3: 'b'
- last['b'] = 1 (seen before at position 1)
- Need to partition at max(2, last['b']) = 2
- Partition: [2,2] (position 2)
- Update last['b'] = 3
- Start new partition at 3

Continue this process...
Final partitions: [9,7,8] ✓"
```

#### Phase 6: Intuition Validation
- **Why greedy works**: Earlier partitions allow more total partitions
- **Why last occurrence**: Determines optimal partition boundary
- **Why O(N)**: Single pass through string
- **Why O(1) space**: Fixed array for 26 letters

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all partitions?"** → Exponential, greedy is optimal
2. **"Should I use sliding window?"** → More complex than needed
3. **"What about Unicode?"** → Clarify character set constraints
4. **"Can I optimize further?"** → Greedy is already optimal

### Real-World Analogy
**Like organizing books by first letter in a library:**
- You have a shelf of books (string) to organize
- Each section must have books with unique first letters
- When you find a book whose first letter appeared before, start new section
- Track where each letter first appeared to know section boundaries
- Goal: maximize number of sections

### Human-Readable Pseudocode
```
function partitionLabels(s):
    if s is empty:
        return []
    
    last = array of size 26 filled with -1
    result = []
    start = 0
    
    for i from 0 to len(s)-1:
        char = s[i] - 'a'
        if last[char] != -1:
            start = max(start, last[char] + 1)
        
        last[char] = i
        
        if i == len(s)-1 or s[i+1] not in current partition:
            result.append(i - start + 1)
    
    return result
```

### Execution Visualization

### Example: s = "ababcbacadefegdehijhklij"
```
Partition Process:
Position: 0, char: 'a', last['a'] = -1, start = 0
Position: 1, char: 'b', last['b'] = -1, start = 0
Position: 2, char: 'a', last['a'] = 0, start = max(0, 0+1) = 1
- Partition: [0,1] (length 2)

Position: 3, char: 'b', last['b'] = 1, start = max(1, 1+1) = 2
- Partition: [2,2] (length 1)

Position: 4, char: 'c', last['c'] = -1, start = 3
Position: 5, char: 'b', last['b'] = 3, start = max(3, 3+1) = 4
- Partition: [3,4] (length 2)

Continue...
Final partitions: [9,7,8] ✓
```

### Key Visualization Points:
- **Character Tracking**: Last occurrence determines partition boundary
- **Greedy Partition**: Create partitions as early as possible
- **Unique Characters**: Each partition contains unique characters
- **Maximum Partitions**: Greedy strategy maximizes partition count

### Memory Layout Visualization:
```
Last Occurrence Tracking:
s = "ababcbacadefegdehijhklij"

Position tracking:
a: [0, 2, 4, 6, 8]
b: [1, 3, 5]
c: [4, 7, 9]
d: [10, 14]
e: [11, 15]
f: [12]
g: [13]
h: [16, 22]
i: [17, 23]
j: [18, 24]
k: [19]
l: [20, 25]

Partition boundaries at positions where characters repeat.
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) time complexity
- **Character Array**: O(1) space for 26 letters
- **Partition Creation**: O(1) time per character
- **Total**: O(N) time, O(1) space

### Alternative Approaches:

#### 1. Sliding Window with Set (O(N²) time, O(N) space)
```go
func partitionLabelsSliding(s string) []int {
    if s == "" {
        return []int{}
    }
    
    result := []int{}
    for i := 0; i < len(s); i++ {
        seen := make(map[byte]bool)
        for j := i; j < len(s); j++ {
            if seen[s[j]] {
                result = append(result, j-i)
                i = j - 1
                break
            }
            seen[s[j]] = true
        }
        
        if len(seen) == len(s)-i {
            result = append(result, len(s)-i)
            break
        }
    }
    
    return result
}
```
- **Pros**: Intuitive approach
- **Cons**: O(N²) time, unnecessary complexity

#### 2. Two-Pass Approach (O(N) time, O(1) space)
```go
func partitionLabelsTwoPass(s string) []int {
    if s == "" {
        return []int{}
    }
    
    // First pass: find all partition boundaries
    last := make([]int, 26)
    for i := range last {
        last[i] = -1
    }
    
    boundaries := make([]bool, len(s))
    start := 0
    
    for i := 0; i < len(s); i++ {
        char := s[i] - 'a'
        if last[char] != -1 {
            start = max(start, last[char] + 1)
        }
        last[char] = i
        boundaries[i] = (i == len(s)-1) || (i+1 < len(s) && last[s[i+1]-'a'] < start)
    }
    
    // Second pass: extract partition lengths
    result := []int{}
    for i := 0; i < len(s); i++ {
        if boundaries[i] {
            result = append(result, i-start+1)
            start = i + 1
        }
    }
    
    return result
}
```
- **Pros**: Same complexity as single pass
- **Cons**: More complex implementation

#### 3. Recursive with Memoization (O(N²) time, O(N) space)
```go
func partitionLabelsRecursive(s string) []int {
    memo := make(map[string][]int)
    return partitionHelper(s, 0, memo)
}

func partitionHelper(s string, start int, memo map[string][]int) []int {
    if start >= len(s) {
        return []int{}
    }
    
    key := fmt.Sprintf("%d", start)
    if val, exists := memo[key]; exists {
        return val
    }
    
    seen := make(map[byte]bool)
    for i := start; i < len(s); i++ {
        if seen[s[i]] {
            result := append([]int{i-start}, partitionHelper(s, i, memo)...)
            memo[key] = result
            return result
        }
        seen[s[i]] = true
    }
    
    result := []int{len(s) - start}
    memo[key] = result
    return result
}
```
- **Pros**: Intuitive approach
- **Cons**: O(N²) time, unnecessary complexity

### Extensions for Interviews:
- **Different Character Sets**: Unicode, case sensitivity
- **Partition Constraints**: Maximum/minimum partition sizes
- **Multiple Strings**: Apply same logic to multiple strings
- **Character Weights**: Different importance for characters
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []string{
		"ababcbacadefegdehijhklij",
		"eccbbbbdec",
		"abac",
		"a",
		"aaaaa",
		"abcde",
		"abacaba",
		"abcdefghijklmnopqrstuvwxyz",
		"zzzzzzzzzz",
		"ababababab",
	}
	
	for i, s := range testCases {
		result := partitionLabels(s)
		fmt.Printf("Test Case %d: \"%s\" -> Partitions: %v\n", i+1, s, result)
	}
}
