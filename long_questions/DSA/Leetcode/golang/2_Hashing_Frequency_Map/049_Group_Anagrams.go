package main

import (
	"fmt"
	"sort"
)

// 49. Group Anagrams
// Time: O(N*K*logK), Space: O(N*K)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	
	for _, str := range strs {
		// Sort the string to create a key
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	
	// Convert map values to slice
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	
	return result
}

// Helper function to sort a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Alternative solution using character count as key (more efficient)
func groupAnagramsOptimized(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	
	for _, str := range strs {
		// Create key based on character count
		key := createCountKey(str)
		anagramMap[key] = append(anagramMap[key], str)
	}
	
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	
	return result
}

// Create key based on character count (26 lowercase letters)
func createCountKey(s string) string {
	count := make([]int, 26)
	for _, char := range s {
		count[char-'a']++
	}
	
	key := fmt.Sprintf("%v", count)
	return key
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hashing with Canonical Representation
- **Canonical Key**: Transform each string to a standard form
- **Grouping Strategy**: Use hash map to group by canonical form
- **Anagram Detection**: Strings with same canonical form are anagrams
- **Frequency Counting**: Character frequency as unique identifier

## 2. PROBLEM CHARACTERISTICS
- **String Grouping**: Group strings that are anagrams of each other
- **Canonical Form**: Each anagram group has unique representation
- **Hash Map Usage**: Efficient grouping using hash map
- **Multiple Solutions**: Both sorting and frequency counting approaches

## 3. SIMILAR PROBLEMS
- Valid Anagram (LeetCode 242) - Check if two strings are anagrams
- Find All Anagrams in a String (LeetCode 438) - Find anagrams of pattern
- Group Shifted Strings (LeetCode 249) - Similar grouping concept
- Palindrome Pairs (LeetCode 336) - String grouping with different criteria

## 4. KEY OBSERVATIONS
- **Anagram property**: Anagrams have same character frequencies
- **Sorting approach**: Sorted strings are identical for anagrams
- **Frequency approach**: Character counts are identical for anagrams
- **Hash map natural fit**: Perfect for grouping by canonical form

## 5. VARIATIONS & EXTENSIONS
- **Unicode characters**: Handle full Unicode character set
- **Case-insensitive**: Treat uppercase and lowercase as same
- **Custom sorting**: Sort groups by size or lexicographically
- **Large datasets**: Optimize for memory efficiency

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are strings lowercase only? What about empty strings?"
- Edge cases: empty array, single string, all same, all different
- Time complexity: O(N*K*logK) sorting, O(N*K) frequency counting
- Space complexity: O(N*K) for storing groups

## 7. COMMON MISTAKES
- Not handling empty strings properly
- Using inefficient string concatenation for keys
- Not considering character set limitations
- Forgetting to handle Unicode characters
- Not optimizing for large inputs

## 8. OPTIMIZATION STRATEGIES
- **Frequency counting**: O(N*K) vs O(N*K*logK) for sorting
- **Character array keys**: More efficient than string concatenation
- **Pre-allocation**: Estimate result size to reduce allocations
- **Early validation**: Check for edge cases upfront

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing books by their letters:**
- You have a pile of books (strings) with different titles
- Books that are anagrams have the same letters, just rearranged
- You want to group books that have identical letter sets
- For each book, you create a "signature" based on its letters
- Books with the same signature belong to the same group

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of strings
2. **Goal**: Group strings that are anagrams of each other
3. **Output**: Array of arrays, each containing anagrams
4. **Constraint**: Order within groups doesn't matter

#### Phase 2: Key Insight Recognition
- **"Canonical representation"** → Need standard form for each anagram group
- **"Hash map grouping"** → Perfect use case for hash map
- **"Two approaches"** → Sorting vs frequency counting
- **"Character analysis"** → Anagrams have identical character counts

#### Phase 3: Strategy Development
```
Human thought process:
"I need to group strings that are anagrams.
For each string, I'll create a canonical key that's the same for all its anagrams.
I'll use a hash map to group strings by this key.
Two ways to create the key:
1. Sort the string (sorted form is same for anagrams)
2. Count character frequencies (frequency array is same for anagrams)
Then I'll collect all groups from the hash map."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty result
- **Empty strings**: Group empty strings together
- **Single string**: Return [[string]]
- **All same strings**: Single group with all strings

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: ["eat", "tea", "tan", "ate", "nat", "bat"]

Human thinking (sorting approach):
"I'll process each string and create sorted keys:

"eat" → sort → "aet" → map["aet"] = ["eat"]
"tea" → sort → "aet" → map["aet"] = ["eat", "tea"]
"tan" → sort → "ant" → map["ant"] = ["tan"]
"ate" → sort → "aet" → map["aet"] = ["eat", "tea", "ate"]
"nat" → sort → "ant" → map["ant"] = ["tan", "nat"]
"bat" → sort → "abt" → map["abt"] = ["bat"]

Final groups:
["eat", "tea", "ate"], ["tan", "nat"], ["bat"]"

Human thinking (frequency approach):
"eat" → count [a:1, e:1, t:1] → key → map[key] = ["eat"]
tea" → count [a:1, e:1, t:1] → same key → map[key] = ["eat", "tea"]
tan" → count [a:1, n:1, t:1] → key → map[key] = ["tan"]
ate" → count [a:1, e:1, t:1] → same key → map[key] = ["eat", "tea", "ate"]
nat" → count [a:1, n:1, t:1] → same key → map[key] = ["tan", "nat"]
bat" → count [a:1, b:1, t:1] → key → map[key] = ["bat"]

Same final result!"
```

#### Phase 6: Intuition Validation
- **Why sorting works**: Anagrams have same sorted representation
- **Why frequency works**: Anagrams have identical character counts
- **Why hash map works**: Efficient grouping by canonical key
- **Why O(N*K)**: N strings, each takes O(K) to process

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compare all pairs?"** → O(N²) time, too slow for large inputs
2. **"Should I use sorting or counting?"** → Counting is faster but more complex
3. **"What about Unicode?"** → Need to handle full character set
4. **"Can I optimize space?"** → Use character arrays instead of strings for keys

### Real-World Analogy
**Like organizing files by their content:**
- You have documents with different text content
- Documents that are anagrams have same letters but different order
- You want to group documents that contain identical letters
- For each document, you create a "fingerprint" based on its letters
- Documents with same fingerprint belong to same group

### Human-Readable Pseudocode
```
function groupAnagrams(strs):
    anagramMap = map()
    
    for str in strs:
        // Create canonical key
        sortedStr = sortString(str)
        
        // Group by canonical key
        if sortedStr not in anagramMap:
            anagramMap[sortedStr] = []
        anagramMap[sortedStr].append(str)
    
    return all values from anagramMap
```

### Execution Visualization

### Example: ["eat", "tea", "tan", "ate", "nat", "bat"]
```
Hash Map Evolution:
Processing "eat":
- sorted("eat") = "aet"
- map = {"aet": ["eat"]}

Processing "tea":
- sorted("tea") = "aet"
- map = {"aet": ["eat", "tea"]}

Processing "tan":
- sorted("tan") = "ant"
- map = {"aet": ["eat", "tea"], "ant": ["tan"]}

Processing "ate":
- sorted("ate") = "aet"
- map = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}

Processing "nat":
- sorted("nat") = "ant"
- map = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}

Processing "bat":
- sorted("bat") = "abt"
- map = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}

Final result: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

### Key Visualization Points:
- **Canonical key creation**: Each string transformed to standard form
- **Hash map grouping**: Strings with same key grouped together
- **Result collection**: Extract all groups from hash map
- **Multiple approaches**: Sorting vs frequency counting

### Memory Layout Visualization:
```
Hash Map Structure:
{
  "aet": ["eat", "tea", "ate"],
  "ant": ["tan", "nat"],
  "abt": ["bat"]
}

Frequency Key for "eat":
[a:1, e:1, t:1] → "[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]"

Frequency Key for "tan":
[a:1, n:1, t:1] → "[1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]"
```

### Time Complexity Breakdown:
- **Sorting approach**: O(N × K × log K) time, O(N × K) space
- **Frequency approach**: O(N × K) time, O(N × K) space
- **N**: number of strings, K: average string length
- **Space**: Hash map storage plus keys

### Alternative Approaches:

#### 1. Frequency Counting (O(N × K) time, O(N × K) space)
```go
func groupAnagramsOptimized(strs []string) [][]string {
    anagramMap := make(map[string][]string)
    
    for _, str := range strs {
        key := createCountKey(str)
        anagramMap[key] = append(anagramMap[key], str)
    }
    
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }
    
    return result
}

func createCountKey(s string) string {
    count := make([]int, 26)
    for _, char := range s {
        count[char-'a']++
    }
    
    key := fmt.Sprintf("%v", count)
    return key
}
```
- **Pros**: Faster than sorting, O(N × K) time
- **Cons**: More complex key generation

#### 2. Prime Multiplication (O(N × K) time, O(N × K) space)
```go
func groupAnagramsPrime(strs []string) [][]string {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
    
    anagramMap := make(map[int64][]string)
    
    for _, str := range strs {
        var key int64 = 1
        for _, char := range str {
            key *= int64(primes[char-'a'])
        }
        anagramMap[key] = append(anagramMap[key], str)
    }
    
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }
    
    return result
}
```
- **Pros**: Unique keys, no collisions
- **Cons**: Limited to lowercase letters, potential overflow

#### 3. Custom Hash Function (O(N × K) time, O(N × K) space)
```go
func groupAnagramsCustomHash(strs []string) [][]string {
    anagramMap := make(map[uint32][]string)
    
    for _, str := range strs {
        var hash uint32 = 0
        count := make([]byte, 26)
        
        for _, char := range str {
            count[char-'a']++
        }
        
        // Custom hash for character count
        for i, c := range count {
            hash = hash*31 + uint32(c)
            hash = hash*31 + uint32(i)
        }
        
        anagramMap[hash] = append(anagramMap[hash], str)
    }
    
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }
    
    return result
}
```
- **Pros**: Fast hash computation
- **Cons**: Potential hash collisions (rare)

### Extensions for Interviews:
- **Unicode Support**: Handle full Unicode character set
- **Case Insensitive**: Treat uppercase and lowercase as same
- **Large Dataset**: Optimize for memory efficiency
- **Streaming**: Process strings as they arrive
- **Custom Sorting**: Sort groups by size or lexicographically
*/
func main() {
	// Test cases
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"abc", "bca", "cab", "def", "fed", "ghi"},
		{"", "", ""},
		{"a", "b", "c", "a", "b", "c"},
	}
	
	for i, strs := range testCases {
		result := groupAnagrams(strs)
		fmt.Printf("Test Case %d: %v\n", i+1, strs)
		fmt.Printf("Grouped Anagrams: %v\n\n", result)
	}
}
