package main

import "fmt"

// Pattern: Hashing / Frequency Map
// Difficulty: Beginner
// Key Concept: Using a Map (HashTable) to store data for O(1) lookups.

/*
INTUITION:
Imagine you are a teacher taking attendance.
Method 1 (Array/List Search): For every student name called, you look through the entire class list one by one to check it off. Slow! (O(N)).
Method 2 (Hash Map): You have a direct lookup sheet. "Alice?" -> Row 1, Seat A. Check. "Bob?" -> Row 2, Seat B. Check. Instant! (O(1)).

In DSA, whenever we need to count occurrences ("Frequency") or check existence ("Has this been seen?"), we use a Map.

PROBLEM:
Find the "First Non-Repeating Character" in a string.
Example: "aabbcdd" -> 'c'

ALGORITHM:
1. Pass 1: "Build the Frequency Map".
   - Iterate through the string.
   - Store the count of every character.
   - 'a': 2, 'b': 2, 'c': 1, 'd': 2
2. Pass 2: "Find the Answer".
   - Iterate through the string AGAIN (in order).
   - Check the map for each character.
   - The first time we see a character with Count == 1, that's our answer.
*/

func firstUniqChar(s string) string {
	// Step 1: Create the map
	// key: rune (character), value: int (count)
	freqMap := make(map[rune]int)

	// Step 2: Populate the map (Pass 1)
	// String: "leetcode"
	// 'l':1, 'e':1
	// 'e':2
	// 't':1 ...
	for _, char := range s {
		freqMap[char]++
	}

	// Step 3: Check distinct counts (Pass 2)
	// We iterate the STRING, not the map, because map order is random in Go.
	// We need the *first* non-repeating char in the original sequence.
	for i, char := range s {
		if freqMap[char] == 1 {
			// Found it!
			// LINE-BY-LINE DRY RUN for "leetcode":
			// i=0, char='l'. Map['l'] == 1? YES. Return "l".
			return string(char)
		}
		// If 'l' was repeated (e.g., "lleetcode"), we would skip and check 'e'.
		_ = i // clear unused var warning if strictly needed
	}

	return "" // If nothing found
}

func main() {
	input := "loveleetcode"
	fmt.Printf("Input: %s\n", input)

	result := firstUniqChar(input)

	if result != "" {
		fmt.Printf("First Non-Repeating Char: %s\n", result)
	} else {
		fmt.Println("No unique character found.")
	}

	/*
		Deep Understanding Note:
		Why 2 passes?
		- Pass 1 builds the "Total Knowledge" of the world (counts).
		- Pass 2 uses that knowledge to make a decision based on order.
		- Time Complexity: O(N) + O(N) = O(2N) -> O(N).
		- Space Complexity: O(U) where U is unique characters (max 26 for lowercase English).
	*/
}
