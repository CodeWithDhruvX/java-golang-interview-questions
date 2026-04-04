package main

import "fmt"

// TrieNode represents a node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// WordDictionary represents the word dictionary with add and search functionality
type WordDictionary struct {
	root *TrieNode
}

// Constructor creates a new WordDictionary
func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// AddWord adds a word to the dictionary
// Time: O(L), Space: O(L) where L is word length
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	
	node.isEnd = true
}

// Search returns true if there is any word in the dictionary that matches the given word
// Time: O(L * 26^W) in worst case where W is number of wildcards, Space: O(L)
func (this *WordDictionary) Search(word string) bool {
	return this.searchHelper(this.root, word, 0)
}

func (this *WordDictionary) searchHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		return node.isEnd
	}
	
	char := rune(word[index])
	
	if char == '.' {
		// Wildcard: try all possible children
		for _, child := range node.children {
			if this.searchHelper(child, word, index+1) {
				return true
			}
		}
		return false
	} else {
		// Normal character: follow the specific path
		child, exists := node.children[char]
		if !exists {
			return false
		}
		return this.searchHelper(child, word, index+1)
	}
}

// Alternative implementation using iterative approach
func (this *WordDictionary) SearchIterative(word string) bool {
	stack := []struct {
		node  *TrieNode
		index int
	}{{this.root, 0}}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if current.index == len(word) {
			if current.node.isEnd {
				return true
			}
			continue
		}
		
		char := rune(word[current.index])
		
		if char == '.' {
			// Add all children to stack
			for _, child := range current.node.children {
				stack = append(stack, struct {
					node  *TrieNode
					index int
				}{child, current.index + 1})
			}
		} else {
			// Add specific child to stack
			child, exists := current.node.children[char]
			if exists {
				stack = append(stack, struct {
					node  *TrieNode
					index int
				}{child, current.index + 1})
			}
		}
	}
	
	return false
}

// Helper method to count words with given pattern
func (this *WordDictionary) CountWords(word string) int {
	return this.countWordsHelper(this.root, word, 0)
}

func (this *WordDictionary) countWordsHelper(node *TrieNode, word string, index int) int {
	if index == len(word) {
		if node.isEnd {
			return 1
		}
		return 0
	}
	
	char := rune(word[index])
	count := 0
	
	if char == '.' {
		// Wildcard: sum counts from all children
		for _, child := range node.children {
			count += this.countWordsHelper(child, word, index+1)
		}
	} else {
		// Normal character: follow specific path
		child, exists := node.children[char]
		if exists {
			count += this.countWordsHelper(child, word, index+1)
		}
	}
	
	return count
}

// Helper method to get all words matching the pattern
func (this *WordDictionary) GetAllMatchingWords(word string) []string {
	var result []string
	this.getAllMatchingHelper(this.root, word, 0, "", &result)
	return result
}

func (this *WordDictionary) getAllMatchingHelper(node *TrieNode, word string, index int, current string, result *[]string) {
	if index == len(word) {
		if node.isEnd {
			*result = append(*result, current)
		}
		return
	}
	
	char := rune(word[index])
	
	if char == '.' {
		// Wildcard: try all children
		for childChar, child := range node.children {
			this.getAllMatchingHelper(child, word, index+1, current+string(childChar), result)
		}
	} else {
		// Normal character: follow specific path
		child, exists := node.children[char]
		if exists {
			this.getAllMatchingHelper(child, word, index+1, current+string(char), result)
		}
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Trie with Wildcard Support
- **Trie Structure**: Standard prefix tree for word storage
- **Wildcard Handling**: '.' character matches any single character
- **Recursive Search**: Backtrack when encountering wildcards
- **Branch Exploration**: Try all possible paths for wildcards

## 2. PROBLEM CHARACTERISTICS
- **Pattern Matching**: Support '.' wildcard for any character
- **Dynamic Word Set**: Add words and search with patterns
- **Backtracking Required**: Explore multiple paths for wildcards
- **Flexible Matching**: Partial pattern matching capability

## 3. SIMILAR PROBLEMS
- Implement Trie (LeetCode 208) - Standard trie without wildcards
- Word Search II (LeetCode 212) - 2D grid with trie patterns
- Implement Magic Dictionary (LeetCode 676) - Trie with character substitution
- Regular Expression Matching (LeetCode 10) - Pattern matching with wildcards

## 4. KEY OBSERVATIONS
- **Wildcard Impact**: '.' forces exploration of all children
- **Recursive Nature**: Natural fit for recursive backtracking
- **Path Branching**: Each wildcard creates multiple search paths
- **Early Termination**: Stop when pattern doesn't match

## 5. VARIATIONS & EXTENSIONS
- **Multiple Wildcards**: Support different wildcard types
- **Character Classes**: Support [a-z] style patterns
- **Quantifiers**: Support * for zero or more matches
- **Case Insensitivity**: Handle uppercase/lowercase matching

## 6. INTERVIEW INSIGHTS
- Always clarify: "What wildcard types? Case sensitivity? Performance requirements?"
- Edge cases: empty strings, all wildcards, no matches
- Time complexity: O(L * 26^W) where W is number of wildcards
- Space complexity: O(total characters) for trie storage
- Key insight: recursive backtracking for wildcard exploration

## 7. COMMON MISTAKES
- Not handling empty strings correctly
- Wrong wildcard logic (treating '.' as literal)
- Infinite recursion with circular references
- Not exploring all children for wildcards
- Stack overflow with deep recursion

## 8. OPTIMIZATION STRATEGIES
- **Recursive Backtracking**: Natural approach for wildcards
- **Iterative Stack**: Avoid recursion depth issues
- **Early Pruning**: Stop when path impossible
- **Memoization**: Cache results for repeated patterns

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a smart dictionary with pattern matching:**
- You have a trie storing words like a normal dictionary
- When searching, '.' acts like a "wildcard" that can match any letter
- At each wildcard, you need to try all possible branches
- It's like a choose-your-own-adventure book where wildcards create multiple paths

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Collection of words to add, patterns to search
2. **Goal**: Support '.' wildcard matching any single character
3. **Constraint**: Efficient pattern matching with dynamic word set
4. **Output**: Boolean indicating if pattern matches any word

#### Phase 2: Key Insight Recognition
- **"Trie natural fit"** → Hierarchical character organization
- **"Wildcard branching"** → '.' requires exploring all children
- **"Recursive backtracking"** → Natural way to handle multiple paths
- **"Path exploration"** → Each wildcard creates search tree

#### Phase 3: Strategy Development
```
Human thought process:
"I need a trie that supports wildcard matching.
Standard trie for word insertion, but search needs special handling:
- Normal character: follow specific child path
- Wildcard '.': try ALL possible children recursively
- Use backtracking to explore all wildcard possibilities
- Return true if ANY path leads to a complete word"

Algorithm:
search(node, pattern, index):
    if index == len(pattern):
        return node.isEnd
    
    char = pattern[index]
    if char == '.':
        for each child in node.children:
            if search(child, pattern, index+1):
                return true
        return false
    else:
        child = node.children[char]
        if child exists:
            return search(child, pattern, index+1)
        return false
```

#### Phase 4: Edge Case Handling
- **Empty pattern**: Check if root.isEnd (empty word case)
- **All wildcards**: Explore exponential number of paths
- **No matches**: Return false after exploring all paths
- **Single character**: Handle edge case efficiently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Words: ["bad", "dad", "mad"]
Pattern: ".ad"

Human thinking:
"I'll search for ".ad" in the trie:

At root, see '.' (wildcard):
- Try 'b' child: search("ad", 'b' node, index=1)
  - Try 'a' child: search("d", 'a' node, index=2)
    - Try 'd' child: search("", 'd' node, index=3)
      - At end, check isEnd: "bad" exists! ✓
      - Return true
    - Return true
  - Return true
- Found match, return true ✓

Pattern "b..":
At root, see 'b' (specific):
- Follow 'b' child: search("..", 'b' node, index=1)
  - See '.' (wildcard): try all children of 'b' node
    - Try 'a' child: search(".", 'a' node, index=2)
      - See '.' (wildcard): try all children of 'a' node
        - Try 'd' child: search("", 'd' node, index=3)
          - At end, check isEnd: "bad" exists! ✓
          - Return true
        - Return true
      - Return true
    - Return true
  - Return true
- Found match, return true ✓"
```

#### Phase 6: Intuition Validation
- **Why trie with wildcards works**: Natural hierarchical organization with branching
- **Why recursion works**: Each wildcard creates multiple search paths
- **Why backtracking needed**: Need to explore all wildcard possibilities
- **Why exponential complexity**: Each wildcard multiplies search paths

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use regex?"** → Problem requires custom implementation
2. **"Should I use iterative approach?"** → Works but more complex
3. **"What about performance?"** → Wildcards make it inherently exponential
4. **"Can I optimize further?"** → Memoization for repeated subproblems
5. **"What about multiple wildcards?"** → Exponential blowup is expected

### Real-World Analogy
**Like a smart phone contact search with wildcards:**
- You have contacts stored in a hierarchical structure
- When you type "J**n", the system searches for:
  - "John" (J-o-h-n)
  - "Joan" (J-o-a-n)  
  - "Jain" (J-a-i-n)
- Each '*' or '.' creates multiple search branches
- System returns true if ANY branch matches a contact

### Human-Readable Pseudocode
```
class WordDictionary:
    root = TrieNode
    
    function addWord(word):
        node = root
        for char in word:
            if char not in node.children:
                node.children[char] = new TrieNode()
            node = node.children[char]
        node.isEnd = true
    
    function search(word):
        return searchHelper(root, word, 0)
    
    function searchHelper(node, word, index):
        if index == len(word):
            return node.isEnd
        
        char = word[index]
        if char == '.':
            for child in node.children:
                if searchHelper(child, word, index + 1):
                    return true
            return false
        else:
            child = node.children[char]
            if child exists:
                return searchHelper(child, word, index + 1)
            return false
```

### Execution Visualization

### Example: Words = ["bad", "dad", "mad"], Pattern = ".ad"
```
Wildcard Search Process:
Pattern: . a d (indices 0,1,2)

Step 0 (index=0, char='.'):
At root, try all children:
- Try 'b': search("ad", 'b' node, index=1)
- Try 'd': search("ad", 'd' node, index=1) 
- Try 'm': search("ad", 'm' node, index=1)

Step 1 (following 'b' path, index=1, char='a'):
At 'b' node, follow 'a' child:
- search("d", 'a' node, index=2)

Step 2 (at 'a' node, index=2, char='d'):
At 'a' node, follow 'd' child:
- search("", 'd' node, index=3)

Step 3 (at 'd' node, index=3):
At end of pattern, check isEnd:
- 'd' node.isEnd = true (word "bad" exists) ✓
- Return true

Multiple paths explored, found match ✓
```

### Key Visualization Points:
- **Wildcard Branching**: '.' creates multiple search paths
- **Recursive Exploration**: Each path explored independently
- **Early Success**: Return true when any path succeeds
- **Complete Search**: Explore all paths if needed

### Memory Layout Visualization:
```
Trie Structure for ["bad", "dad", "mad"]:
root
 ├── 'b'
 │    └── 'a'
 │         └── 'd' (END) ← "bad"
 ├── 'd'
 │    └── 'a'
 │         └── 'd' (END) ← "dad"
 └── 'm'
      └── 'a'
           └── 'd' (END) ← "mad"

Search ".ad":
At root '.' → explore 'b', 'd', 'm' paths
'b' path → 'a' → 'd' (END) ✓ Match found!
```

### Time Complexity Breakdown:
- **AddWord**: O(L) time, O(L) space where L is word length
- **Search**: O(L * 26^W) time where W is number of wildcards
- **Space**: O(total characters) for trie storage
- **Worst Case**: Exponential when pattern has many wildcards

### Alternative Approaches:

#### 1. Hash Set with Pattern Matching (O(N * L) time, O(N * L) space)
```go
type WordDictionary struct {
    words []string
}

func (wd *WordDictionary) AddWord(word string) {
    wd.words = append(wd.words, word)
}

func (wd *WordDictionary) Search(word string) bool {
    for _, dictWord := range wd.words {
        if len(dictWord) != len(word) {
            continue
        }
        
        match := true
        for i := 0; i < len(word); i++ {
            if word[i] != '.' && word[i] != dictWord[i] {
                match = false
                break
            }
        }
        
        if match {
            return true
        }
    }
    return false
}
```
- **Pros**: Simple to implement
- **Cons**: O(N) search time, no prefix optimization

#### 2. Iterative Stack Approach (O(L * 26^W) time, O(L) space)
```go
func (wd *WordDictionary) SearchIterative(word string) bool {
    stack := []struct {
        node  *TrieNode
        index int
    }{{wd.root, 0}}
    
    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if current.index == len(word) {
            if current.node.isEnd {
                return true
            }
            continue
        }
        
        char := rune(word[current.index])
        if char == '.' {
            for _, child := range current.node.children {
                stack = append(stack, struct {
                    node  *TrieNode
                    index int
                }{child, current.index + 1})
            }
        } else {
            child, exists := current.node.children[char]
            if exists {
                stack = append(stack, struct {
                    node  *TrieNode
                    index int
                }{child, current.index + 1})
            }
        }
    }
    
    return false
}
```
- **Pros**: Avoids recursion depth issues
- **Cons**: More complex implementation

#### 3. Memoized Search (O(L * 26^W) time, O(N * L) space)
```go
type WordDictionary struct {
    root    *TrieNode
    cache   map[string]bool
}

func (wd *WordDictionary) Search(word string) bool {
    if result, exists := wd.cache[word]; exists {
        return result
    }
    
    result := wd.searchHelper(wd.root, word, 0)
    wd.cache[word] = result
    return result
}
```
- **Pros**: Cache repeated pattern searches
- **Cons**: Additional memory overhead

### Extensions for Interviews:
- **Multiple Wildcards**: Support '*', '+', '?' patterns
- **Character Classes**: Support [a-z], [0-9] patterns
- **Quantifiers**: Support repetition patterns
- **Case Insensitivity**: Handle uppercase/lowercase matching
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	wd := ConstructorWordDictionary()
	
	fmt.Println("=== Testing AddWord and Search ===")
	
	// Add words
	words := []string{"bad", "dad", "mad", "baddie", "daddy", "apple"}
	for _, word := range words {
		wd.AddWord(word)
		fmt.Printf("Added: %s\n", word)
	}
	
	// Test search
	searchWords := []string{
		"pad", "bad", ".ad", "b..", "d..y", "b....e", "a..le", "....",
	}
	
	for _, word := range searchWords {
		found := wd.Search(word)
		count := wd.CountWords(word)
		matching := wd.GetAllMatchingWords(word)
		
		fmt.Printf("Search '%s': %t (count: %d, matches: %v)\n", word, found, count, matching)
	}
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	emptyWD := ConstructorWordDictionary()
	fmt.Printf("Empty dictionary search 'test': %t\n", emptyWD.Search("test"))
	
	// Add single character words
	emptyWD.AddWord("a")
	emptyWD.AddWord("b")
	fmt.Printf("After adding 'a', 'b': search '.' = %t\n", emptyWD.Search("."))
	fmt.Printf("Search 'a': %t, search 'b': %t\n", emptyWD.Search("a"), emptyWD.Search("b"))
	
	// Test with multiple wildcards
	fmt.Println("\n=== Testing Multiple Wildcards ===")
	wd.AddWord("test")
	wd.AddWord("tent")
	wd.AddWord("text")
	
	wildcardTests := []string{"t..t", "t.e.", ".e.t", "...."}
	for _, pattern := range wildcardTests {
		matching := wd.GetAllMatchingWords(pattern)
		fmt.Printf("Pattern '%s' matches: %v\n", pattern, matching)
	}
}
