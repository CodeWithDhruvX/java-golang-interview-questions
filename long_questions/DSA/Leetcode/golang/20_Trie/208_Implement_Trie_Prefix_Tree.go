package main

import "fmt"

// TrieNode represents a node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the trie data structure
type Trie struct {
	root *TrieNode
}

// Constructor creates a new Trie
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert inserts a word into the trie
// Time: O(L), Space: O(L) where L is word length
func (this *Trie) Insert(word string) {
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

// Search returns true if the word is in the trie
// Time: O(L), Space: O(1)
func (this *Trie) Search(word string) bool {
	node := this.root
	
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	
	return node.isEnd
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
// Time: O(L), Space: O(1)
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	
	return true
}

// Helper function to visualize the trie structure
func (this *Trie) Visualize() {
	fmt.Println("Trie Structure:")
	this.visualizeHelper(this.root, "", true)
}

func (this *Trie) visualizeHelper(node *TrieNode, prefix string, isLast bool) {
	if node == nil {
		return
	}
	
	// Print current node
	if node.isEnd {
		fmt.Printf("%s└── (END)\n", prefix)
	} else {
		fmt.Printf("%s└── (node)\n", prefix)
	}
	
	// Print children
	count := 0
	for char, child := range node.children {
		count++
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		
		fmt.Printf("%s└── %c\n", newPrefix, char)
		this.visualizeHelper(child, newPrefix+"    ", count == len(node.children))
	}
}

// Additional useful methods

// GetAllWords returns all words in the trie
func (this *Trie) GetAllWords() []string {
	var words []string
	this.getAllWordsHelper(this.root, "", &words)
	return words
}

func (this *Trie) getAllWordsHelper(node *TrieNode, current string, words *[]string) {
	if node.isEnd {
		*words = append(*words, current)
	}
	
	for char, child := range node.children {
		this.getAllWordsHelper(child, current+string(char), words)
	}
}

// CountWords returns the number of words in the trie
func (this *Trie) CountWords() int {
	return this.countWordsHelper(this.root)
}

func (this *Trie) countWordsHelper(node *TrieNode) int {
	count := 0
	if node.isEnd {
		count++
	}
	
	for _, child := range node.children {
		count += this.countWordsHelper(child)
	}
	
	return count
}

// Delete removes a word from the trie
func (this *Trie) Delete(word string) bool {
	return this.deleteHelper(this.root, word, 0)
}

func (this *Trie) deleteHelper(node *TrieNode, word string, depth int) bool {
	if depth == len(word) {
		if !node.isEnd {
			return false // Word doesn't exist
		}
		node.isEnd = false
		return len(node.children) == 0
	}
	
	char := rune(word[depth])
	child, exists := node.children[char]
	if !exists {
		return false // Word doesn't exist
	}
	
	shouldDeleteChild := this.deleteHelper(child, word, depth+1)
	if shouldDeleteChild {
		delete(node.children, char)
		return len(node.children) == 0 && !node.isEnd
	}
	
	return false
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Trie (Prefix Tree) for String Operations
- **Tree Structure**: Each node represents a prefix, children for next characters
- **Prefix Sharing**: Common prefixes share nodes, saving space
- **Character Navigation**: Navigate tree character by character
- **End Markers**: Boolean flags mark word boundaries

## 2. PROBLEM CHARACTERISTICS
- **String Operations**: Insert, search, prefix matching
- **Prefix Optimization**: Share common prefixes efficiently
- **Dynamic Structure**: Support insertions and deletions
- **Alphabet Support**: Handle character sets (ASCII, Unicode)

## 3. SIMILAR PROBLEMS
- Design Add and Search Words (LeetCode 211) - Trie with wildcards
- Word Search II (LeetCode 212) - 2D grid with Trie
- Implement Magic Dictionary (LeetCode 676) - Trie for word transformations
- Concatenated Words (LeetCode 472) - Trie for word combinations

## 4. KEY OBSERVATIONS
- **Prefix Sharing**: Common prefixes share path in tree
- **Space Efficiency**: Shared prefixes reduce memory usage
- **Time Complexity**: O(L) for operations where L is string length
- **Node Structure**: Children map + end marker

## 5. VARIATIONS & EXTENSIONS
- **Compressed Trie**: Radix tree for space optimization
- **Ternary Search Trie**: Three-way branching for search optimization
- **Patricia Trie**: Path compression for long strings
- **Different Alphabets**: Support Unicode, case sensitivity

## 6. INTERVIEW INSIGHTS
- Always clarify: "What character set? Case sensitivity? Deletion needed?"
- Edge cases: empty strings, single characters, duplicate words
- Time complexity: O(L) for operations, O(total characters) space
- Key insight: prefix sharing saves space and enables prefix operations
- Memory vs Time tradeoff: More children vs compression

## 7. COMMON MISTAKES
- Not handling empty strings correctly
- Wrong character type (rune vs byte vs string)
- Not marking end of words properly
- Memory leaks in node creation
- Incorrect deletion logic

## 8. OPTIMIZATION STRATEGIES
- **Standard Trie**: O(L) time, O(total characters) space
- **Array Children**: Use fixed-size array for known alphabets
- **Compressed Trie**: Path compression for space efficiency
- **Lazy Deletion**: Mark nodes as deleted instead of immediate removal

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a dictionary organized by prefixes:**
- You have a tree where each path represents a word prefix
- Common prefixes share the same path (like "app" and "apple")
- Each node knows if it's the end of a complete word
- Navigation follows characters down the tree
- Like a phone's T9 predictive text system

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Collection of words for prefix tree
2. **Goal**: Efficient insert, search, and prefix operations
3. **Constraint**: Optimize for prefix-based operations
4. **Output**: Trie data structure with required operations

#### Phase 2: Key Insight Recognition
- **"Prefix sharing natural fit"** → Common prefixes share nodes
- **"Tree structure"** → Hierarchical character organization
- **"End markers"** → Need to distinguish prefixes from complete words
- **"Character navigation"** → Follow tree path character by character

#### Phase 3: Strategy Development
```
Human thought process:
"I need a data structure for efficient prefix operations.
A trie is perfect:
- Each node represents a prefix
- Children represent next possible characters
- Common prefixes share the same path
- Navigation follows characters down the tree

Operations:
- Insert: Follow/create path for word
- Search: Navigate path, check end marker
- StartsWith: Navigate path, don't need end marker"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Handle as valid word or ignore based on requirements
- **Single character**: Create immediate child of root
- **Duplicate words**: Handle gracefully (already exists)
- **Character set**: Support required alphabet (ASCII, Unicode)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Words: ["apple", "app", "application"]

Human thinking:
"I'll build a trie step by step:

Insert "apple":
- root -> 'a' -> 'p' -> 'p' -> 'l' -> 'e' (END)
- Path: root-a-p-p-l-e, mark 'e' as word end

Insert "app":
- root -> 'a' -> 'p' -> 'p' (END)
- Path: root-a-p-p, mark 'p' as word end
- Note: "app" prefix is shared with "apple"

Insert "application":
- root -> 'a' -> 'p' -> 'p' -> 'l' -> 'i' -> 'c' -> 'a' -> 't' -> 'i' -> 'o' -> 'n' (END)
- Path: root-a-p-p-l-i-c-a-t-i-o-n, mark 'n' as word end
- Note: "app" prefix is shared, "appl" prefix is shared

Final structure has shared prefixes ✓"
```

#### Phase 6: Intuition Validation
- **Why trie works**: Shared prefixes reduce memory and enable prefix operations
- **Why O(L)**: Each operation follows path of length L
- **Why space efficient**: Common prefixes share nodes
- **Why tree structure**: Natural hierarchical organization

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use hash set?"** → No prefix operations, more memory
2. **"Should I use array for children?"** → Faster for known alphabets
3. **"What about case sensitivity?"** → Clarify character handling requirements
4. **"Can I optimize space?"** → Use compressed trie or array children
5. **"What about Unicode?"** → Use rune instead of byte for Go

### Real-World Analogy
**Like a phone's T9 predictive text system:**
- Each key press narrows down possible words
- Common prefixes are shared (like "app" for "apple", "app", "application")
- The system remembers word endings (end markers)
- As you type, it follows the character path
- Can suggest completions based on current prefix

### Human-Readable Pseudocode
```
class Trie:
    root = TrieNode with children map and isEnd flag
    
    function insert(word):
        node = root
        for char in word:
            if char not in node.children:
                node.children[char] = new TrieNode()
            node = node.children[char]
        node.isEnd = true
    
    function search(word):
        node = root
        for char in word:
            if char not in node.children:
                return false
            node = node.children[char]
        return node.isEnd
    
    function startsWith(prefix):
        node = root
        for char in prefix:
            if char not in node.children:
                return false
            node = node.children[char]
        return true
```

### Execution Visualization

### Example: Words = ["apple", "app", "application"]
```
Trie Construction:
Step 1: Insert "apple"
root
 └── 'a'
     └── 'p'
         └── 'p'
             └── 'l'
                 └── 'e' (END)

Step 2: Insert "app" (shares prefix)
root
 └── 'a'
     └── 'p'
         └── 'p' (END) ← New word end marker
             └── 'l'
                 └── 'e' (END)

Step 3: Insert "application" (shares longer prefix)
root
 └── 'a'
     └── 'p'
         └── 'p' (END)
             └── 'l'
                 └── 'e' (END)
                     └── 'i'
                         └── 'c'
                             └── 'a'
                                 └── 't'
                                     └── 'i'
                                         └── 'o'
                                             └── 'n' (END)

Final trie with shared prefixes ✓
```

### Key Visualization Points:
- **Prefix Sharing**: Common prefixes share the same path
- **End Markers**: Distinguish complete words from prefixes
- **Character Navigation**: Follow tree path character by character
- **Space Efficiency**: Shared prefixes reduce memory usage

### Memory Layout Visualization:
```
Trie Node Structure:
TrieNode {
    children: map[rune]*TrieNode
    isEnd: bool
}

Example node for 'p' in "app", "apple":
children = {
    'l': TrieNode{...},  // for "apple"
    'p': TrieNode{isEnd: true}  // for "app"
}
isEnd = false  // "app" ends here, "apple" continues
```

### Time Complexity Breakdown:
- **Insert**: O(L) time, O(L) space where L is word length
- **Search**: O(L) time, O(1) space
- **StartsWith**: O(P) time, O(1) space where P is prefix length
- **Space**: O(total characters) for all inserted words
- **Optimal**: Best for prefix-based operations

### Alternative Approaches:

#### 1. Hash Set (O(L) time, O(N*L) space)
```go
type HashSet struct {
    words map[string]bool
}

func (hs *HashSet) Insert(word string) {
    hs.words[word] = true
}

func (hs *HashSet) Search(word string) bool {
    return hs.words[word]
}

func (hs *HashSet) StartsWith(prefix string) bool {
    for word := range hs.words {
        if strings.HasPrefix(word, prefix) {
            return true
        }
    }
    return false
}
```
- **Pros**: Simple to implement
- **Cons**: O(N) for prefix operations, more memory

#### 2. Sorted Array (O(log N) time, O(N*L) space)
```go
type SortedArray struct {
    words []string
}

func (sa *SortedArray) Insert(word string) {
    // Insert while maintaining sorted order
    // Binary search for insertion point
}

func (sa *SortedArray) Search(word string) bool {
    // Binary search for exact match
}

func (sa *SortedArray) StartsWith(prefix string) bool {
    // Binary search for first word with prefix
    // Check range of words with same prefix
}
```
- **Pros**: Memory efficient for search
- **Cons**: Complex prefix operations, O(log N) insertions

#### 3. Array-Based Trie (O(1) children access, O(Σ) space)
```go
type ArrayTrieNode struct {
    children [26]*ArrayTrieNode  // For lowercase English
    isEnd    bool
}

type ArrayTrie struct {
    root *ArrayTrieNode
}
```
- **Pros**: O(1) child access, no map overhead
- **Cons**: Fixed alphabet size, memory inefficient for sparse usage

### Extensions for Interviews:
- **Compressed Trie**: Radix tree for space optimization
- **Ternary Search Trie**: Three-way branching for search optimization
- **Patricia Trie**: Path compression for long strings
- **Different Alphabets**: Support Unicode, case sensitivity
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	trie := Constructor()
	
	// Test insertion and search
	fmt.Println("=== Testing Insert and Search ===")
	words := []string{"apple", "app", "application", "apt", "bat", "batch"}
	for _, word := range words {
		trie.Insert(word)
		fmt.Printf("Inserted: %s\n", word)
	}
	
	// Test search
	searchWords := []string{"apple", "app", "appl", "bat", "batter"}
	for _, word := range searchWords {
		found := trie.Search(word)
		fmt.Printf("Search %s: %t\n", word, found)
	}
	
	// Test starts with
	fmt.Println("\n=== Testing StartsWith ===")
	prefixes := []string{"app", "ap", "bat", "cat"}
	for _, prefix := range prefixes {
		found := trie.StartsWith(prefix)
		fmt.Printf("StartsWith %s: %t\n", prefix, found)
	}
	
	// Test additional methods
	fmt.Println("\n=== Additional Methods ===")
	fmt.Printf("Total words: %d\n", trie.CountWords())
	
	allWords := trie.GetAllWords()
	fmt.Printf("All words: %v\n", allWords)
	
	// Test deletion
	fmt.Println("\n=== Testing Delete ===")
	trie.Delete("app")
	fmt.Printf("After deleting 'app', Search 'app': %t\n", trie.Search("app"))
	fmt.Printf("After deleting 'app', Search 'apple': %t\n", trie.Search("apple"))
	fmt.Printf("Total words after deletion: %d\n", trie.CountWords())
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	emptyTrie := Constructor()
	fmt.Printf("Empty trie search 'test': %t\n", emptyTrie.Search("test"))
	fmt.Printf("Empty trie startsWith 'test': %t\n", emptyTrie.StartsWith("test"))
	
	emptyTrie.Insert("")
	fmt.Printf("Insert empty string, count: %d\n", emptyTrie.CountWords())
}
