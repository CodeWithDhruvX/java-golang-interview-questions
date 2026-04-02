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
