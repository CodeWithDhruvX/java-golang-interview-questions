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
