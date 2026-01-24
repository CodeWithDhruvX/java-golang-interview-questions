package main

import "fmt"

// Pattern: Trie (Prefix Tree)
// Difficulty: Medium
// Key Concept: Storing strings in a tree structure where nodes share common prefixes to save space and search time.

/*
INTUITION:
Imagine a dictionary.
"Car", "Cat", "Cap".
They all start with "C", then "a".
Instead of storing "Car", "Cat", "Cap" separately (3 strings), we store a path:
Root -> 'C' -> 'a' -> 'r' (End)
                   -> 't' (End)
                   -> 'p' (End)
This "Tree" structure allows us to:
1. Search efficiently (O(Length of Word), not O(Size of Dictionary)).
2. Find all words starting with "Ca" instantly.

PROBLEM:
Implement a Trie (Prefix Tree) with `Insert`, `Search`, and `StartsWith`.

ALGORITHM:
Node Structure:
- `children`: A map or array (size 26) pointing to next letters.
- `isEnd`: Boolean to mark if a word ends here (e.g., distinguishing "apple" from "app").

Insert:
- Start at Root. For each char, check if child exists. If not, create it. Move down. Mark `isEnd = true` at last node.

Search:
- Start at Root. traverse. If path breaks, return False. If we finish word, check `isEnd`.
*/

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		// Does path exist?
		if _, exists := node.children[char]; !exists {
			// Create new branch
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		// Move down
		node = node.children[char]
	}
	// Mark the end of the word
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		if nextNode, exists := node.children[char]; exists {
			node = nextNode
		} else {
			return false // Broken path
		}
	}
	// We matched the path, but is it a whole word?
	// e.g. "app" exists, but we searched "apple"? (Path broken)
	// e.g. "apple" exists, but we searched "app"? (Path valid, but is "app" marked as a word?)
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		if nextNode, exists := node.children[char]; exists {
			node = nextNode
		} else {
			return false
		}
	}
	// Path exists. We don't care if it's a full word or not.
	return true
}

func main() {
	obj := Constructor()
	obj.Insert("apple")

	fmt.Printf("Inserted 'apple'\n")
	fmt.Printf("Search 'apple': %v\n", obj.Search("apple"))     // True
	fmt.Printf("Search 'app': %v\n", obj.Search("app"))         // False (It's a prefix, not marked as word yet)
	fmt.Printf("StartsWith 'app': %v\n", obj.StartsWith("app")) // True

	obj.Insert("app")
	fmt.Printf("Inserted 'app'\n")
	fmt.Printf("Search 'app': %v\n", obj.Search("app")) // True
}
