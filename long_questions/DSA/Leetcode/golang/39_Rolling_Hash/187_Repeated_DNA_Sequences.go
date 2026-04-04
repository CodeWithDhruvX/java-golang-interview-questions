package main

import (
	"fmt"
	"math"
)

// 187. Repeated DNA Sequences - Rabin-Karp with Rolling Hash
// Time: O(N*L) where N is string length, L is pattern length (10), Space: O(N)
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	// Hash set to store seen sequences
	seen := make(map[string]bool)
	repeated := make(map[string]bool)
	
	for i := 0; i <= len(s)-10; i++ {
		sequence := s[i : i+10]
		
		if seen[sequence] {
			repeated[sequence] = true
		} else {
			seen[sequence] = true
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(repeated))
	for sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

// Optimized version using rolling hash
func findRepeatedDnaSequencesRollingHash(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	// Map for character to integer (A=0, C=1, G=2, T=3)
	charMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	
	seen := make(map[int]bool)
	repeated := make(map[int]string)
	
	// Base for hash calculation
	base := 4
	windowSize := 10
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < windowSize; i++ {
		hash = hash*base + charMap[s[i]]
	}
	
	seen[hash] = true
	
	// Rolling hash
	for i := 1; i <= len(s)-windowSize; i++ {
		// Remove leftmost character contribution
		leftChar := charMap[s[i-1]]
		hash -= leftChar * int(math.Pow(float64(base), float64(windowSize-1)))
		
		// Add new character
		hash = hash*base + charMap[s[i+windowSize-1]]
		
		if seen[hash] {
			sequence := s[i : i+windowSize]
			repeated[hash] = sequence
		} else {
			seen[hash] = true
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(repeated))
	for _, sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

// More efficient rolling hash with bit manipulation
func findRepeatedDnaSequencesBitHash(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	// Use 2 bits per character (since we have 4 characters)
	// 10 characters * 2 bits = 20 bits, fits in 32-bit integer
	
	seen := make(map[int]bool)
	repeated := make(map[int]string)
	
	// Bit mask for 20 bits
	mask := (1 << 20) - 1
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < 10; i++ {
		hash = (hash << 2) | charToBits(s[i])
	}
	
	seen[hash] = true
	
	// Rolling hash
	for i := 1; i <= len(s)-10; i++ {
		// Remove leftmost 2 bits and add new 2 bits
		hash = ((hash << 2) & mask) | charToBits(s[i+9])
		
		if seen[hash] {
			sequence := s[i : i+10]
			repeated[hash] = sequence
		} else {
			seen[hash] = true
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(repeated))
	for _, sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

func charToBits(c byte) int {
	switch c {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'G':
		return 2
	case 'T':
		return 3
	default:
		return 0
	}
}

// Alternative using different hash function
func findRepeatedDnaSequencesAlternative(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	// Use a different rolling hash approach
	const L = 10
	const base = 7
	const mod = 1000000007
	
	seen := make(map[int]bool)
	repeated := make(map[int]string)
	
	// Precompute base^L mod
	powL := 1
	for i := 0; i < L; i++ {
		powL = (powL * base) % mod
	}
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < L; i++ {
		hash = (hash*base + int(s[i])) % mod
	}
	
	seen[hash] = true
	
	// Rolling hash
	for i := 1; i <= len(s)-L; i++ {
		// Remove leftmost character
		hash = (hash - int(s[i-1])*powL%mod + mod) % mod
		
		// Add new character
		hash = (hash*base + int(s[i+L-1])) % mod
		
		if seen[hash] {
			sequence := s[i : i+L]
			repeated[hash] = sequence
		} else {
			seen[hash] = true
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(repeated))
	for _, sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

// Version using multiple hash functions to reduce collisions
func findRepeatedDnaSequencesDoubleHash(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	const L = 10
	
	// Two different hash functions
	seen := make(map[[2]int]bool)
	repeated := make(map[[2]int]string)
	
	// First hash: base 4, mod large prime
	hash1 := 0
	base1 := 4
	mod1 := 1000000007
	
	// Second hash: base 7, mod different prime
	hash2 := 0
	base2 := 7
	mod2 := 1000000009
	
	// Precompute powers
	pow1 := 1
	pow2 := 1
	for i := 0; i < L; i++ {
		pow1 = (pow1 * base1) % mod1
		pow2 = (pow2 * base2) % mod2
	}
	
	// Calculate initial hashes
	for i := 0; i < L; i++ {
		val := charToBits(s[i])
		hash1 = (hash1*base1 + val) % mod1
		hash2 = (hash2*base2 + val) % mod2
	}
	
	seen[[2]int{hash1, hash2}] = true
	
	// Rolling hash
	for i := 1; i <= len(s)-L; i++ {
		// Remove leftmost character
		leftVal := charToBits(s[i-1])
		hash1 = (hash1 - leftVal*pow1%mod1 + mod1) % mod1
		hash2 = (hash2 - leftVal*pow2%mod2 + mod2) % mod2
		
		// Add new character
		newVal := charToBits(s[i+L-1])
		hash1 = (hash1*base1 + newVal) % mod1
		hash2 = (hash2*base2 + newVal) % mod2
		
		hashPair := [2]int{hash1, hash2}
		
		if seen[hashPair] {
			sequence := s[i : i+L]
			repeated[hashPair] = sequence
		} else {
			seen[hashPair] = true
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(repeated))
	for _, sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

// Brute force approach for comparison
func findRepeatedDnaSequencesBruteForce(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	
	seen := make(map[string]bool)
	repeated := make(map[string]bool)
	
	for i := 0; i <= len(s)-10; i++ {
		sequence := s[i : i+10]
		
		if seen[sequence] {
			repeated[sequence] = true
		} else {
			seen[sequence] = true
		}
	}
	
	result := make([]string, 0, len(repeated))
	for sequence := range repeated {
		result = append(result, sequence)
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Rolling Hash for Pattern Detection
- **Rabin-Karp Algorithm**: Efficient substring hashing and matching
- **Sliding Window**: Fixed-size window (10 characters) for DNA sequences
- **Hash Collisions**: Handle collisions with verification or multiple hashes
- **Bit Manipulation**: Optimize using 2 bits per DNA character

## 2. PROBLEM CHARACTERISTICS
- **Fixed Pattern Length**: DNA sequences of exactly 10 characters
- **Repeated Detection**: Find sequences appearing at least twice
- **Alphabet Constraint**: Only 4 characters (A, C, G, T)
- **Large Input**: String can be up to 10^5 characters

## 3. SIMILAR PROBLEMS
- Longest Duplicate Substring (LeetCode 1044) - Rolling hash with binary search
- Substring with Concatenation of All Words (LeetCode 30) - Hash-based pattern matching
- Detect Squares (LeetCode 593) - Pattern detection with hashing
- Repeated String Match (LeetCode 686) - Pattern repetition detection

## 4. KEY OBSERVATIONS
- **Fixed Window**: 10-character sequences enable efficient rolling hash
- **Small Alphabet**: 4 characters allow bit-level optimization
- **Hash Efficiency**: O(1) hash updates vs O(10) string comparison
- **Collision Handling**: Multiple hash functions reduce false positives

## 5. VARIATIONS & EXTENSIONS
- **Simple Hash**: Direct string hashing with maps
- **Rolling Hash**: O(1) hash updates for sliding window
- **Bit Hash**: 2-bit encoding for maximum efficiency
- **Double Hash**: Two hash functions for collision reduction

## 6. INTERVIEW INSIGHTS
- Always clarify: "String length constraints? Alphabet size? Collision tolerance?"
- Edge cases: empty string, too short strings, no repeats
- Time complexity: O(N) for optimized rolling hash, O(N×10) for simple approach
- Space complexity: O(N) for hash storage
- Key insight: fixed pattern length enables perfect rolling hash optimization

## 7. COMMON MISTAKES
- Not handling hash collisions properly
- Wrong rolling hash update formula
- Missing edge cases for short strings
- Inefficient character encoding
- Not verifying actual sequences on hash collision

## 8. OPTIMIZATION STRATEGIES
- **Simple Hash**: O(N×10) time, O(N) space - straightforward approach
- **Rolling Hash**: O(N) time, O(N) space - efficient sliding window
- **Bit Hash**: O(N) time, O(N) space - maximum efficiency
- **Double Hash**: O(N) time, O(N) space - collision reduction

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding repeated phrases in a book:**
- Each 10-character DNA sequence is like a 10-word phrase
- You slide a window through the text, checking each phrase
- Instead of comparing phrases character by character, you use a hash
- When you see the same hash, you've found a repeated phrase
- Like finding plagiarism using fingerprinting techniques

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: DNA string with characters A, C, G, T
2. **Goal**: Find all 10-character sequences appearing at least twice
3. **Rules**: Return actual sequences, not just counts
4. **Output**: List of repeated 10-character DNA sequences

#### Phase 2: Key Insight Recognition
- **"Fixed window natural"** → 10-character sequences enable sliding window
- **"Hash optimization possible"** → Avoid O(10) string comparisons
- **"Small alphabet advantage"** → 4 characters allow bit encoding
- **"Rolling hash efficient"** → O(1) updates vs O(10) comparisons

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find repeated 10-character DNA sequences.
Naive approach: compare each 10-char window with all others O(N²×10).

Rolling Hash Approach:
1. Encode DNA characters (A=0, C=1, G=2, T=3)
2. Use sliding window of size 10
3. Calculate rolling hash: hash = hash*base + newChar
4. Remove oldChar contribution: hash -= oldChar*base^9
5. Store hashes in map, track duplicates
6. Return sequences with hash count > 1

This gives O(N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty result
- **Too short**: Return empty if length < 10
- **No repeats**: Return empty result
- **All same**: Handle efficiently with bit encoding

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Human thinking:
"Rolling Hash Approach:
Step 1: Encode characters A=0, C=1, G=2, T=3
Step 2: Initial window "AAAAACCCCC" (positions 0-9)
Hash = 0*4^9 + 0*4^8 + ... + 1*4^0 = 1024

Step 3: Slide to positions 1-10 "AAAACCCCCA"
Remove leftmost 'A' (0*4^9 = 0): hash = 1024 - 0 = 1024
Add new 'A' (0): hash = 1024*4 + 0 = 4096

Step 4: Continue sliding...
When hash repeats, we've found duplicate sequence!

Key sequences found: "AAAAACCCCC", "CCCCCAAAAA" ✓"
```

#### Phase 6: Intuition Validation
- **Why rolling hash**: O(1) updates vs O(10) comparisons
- **Why base 4**: Perfect for 4-character DNA alphabet
- **Why bit encoding**: 2 bits per character = 20 bits total
- **Why O(N)**: Single pass through string with O(1) per window

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compare strings directly?"** → O(N×10) vs O(N) with hash
2. **"Should I use double hash?"** → Yes, to reduce collision probability
3. **"What about bit manipulation?"** → Maximum efficiency for DNA
4. **"Can I use different base?"** → Base 4 is optimal for DNA alphabet
5. **"Why handle collisions?"** → Hash collisions can give false positives

### Real-World Analogy
**Like DNA fingerprinting in forensic science:**
- Each 10-character sequence is like a DNA fingerprint fragment
- You scan through the DNA sample looking for matching fragments
- Instead of comparing entire sequences, you use hash fingerprints
- When fingerprints match, you've found repeated DNA patterns
- Like identifying genetic markers in forensic analysis

### Human-Readable Pseudocode
```
function findRepeatedDnaSequences(s):
    if length(s) < 10:
        return []
    
    # Encode DNA characters
    charMap = {'A': 0, 'C': 1, 'G': 2, 'T': 3}
    seen = hash set
    repeated = hash set
    
    # Rolling hash parameters
    base = 4
    windowSize = 10
    
    # Calculate initial hash
    hash = 0
    for i from 0 to 9:
        hash = hash * base + charMap[s[i]]
    
    seen.add(hash)
    
    # Rolling hash
    for i from 1 to length(s) - 10:
        # Remove leftmost character
        leftChar = charMap[s[i-1]]
        hash -= leftChar * (base ^ 9)
        
        # Add new character
        hash = hash * base + charMap[s[i+9]]
        
        if hash in seen:
            repeated.add(s[i:i+10])
        else:
            seen.add(hash)
    
    return list(repeated)
```

### Execution Visualization

### Example: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
```
Rolling Hash Process:

Character encoding: A=0, C=1, G=2, T=3

Window positions 0-9: "AAAAACCCCC"
Hash = 0*4^9 + 0*4^8 + 0*4^7 + 0*4^6 + 0*4^5 + 1*4^4 + 1*4^3 + 1*4^2 + 1*4^1 + 1*4^0
Hash = 0 + 0 + 0 + 0 + 0 + 256 + 64 + 16 + 4 + 1 = 341

Slide to positions 1-10: "AAAACCCCCA"
Remove leftmost 'A' (0*4^9): hash = 341 - 0 = 341
Add new 'A' (0): hash = 341*4 + 0 = 1364

Continue sliding...
When hash 341 appears again at positions 10-19: "AAAAACCCCC"
Found duplicate! ✓

Key repeated sequences:
- "AAAAACCCCC" (appears at positions 0 and 10)
- "CCCCCAAAAA" (appears at positions 5 and 15)
```

### Key Visualization Points:
- **Sliding Window**: Fixed 10-character window moves through string
- **Hash Update**: Remove leftmost, add rightmost in O(1)
- **Collision Detection**: Hash values identify potential duplicates
- **Sequence Storage**: Store actual sequences for final result

### DNA Encoding Visualization:
```
Character → Bits → Decimal
A → 00 → 0
C → 01 → 1  
G → 10 → 2
T → 11 → 3

Sequence "ACGT" → 00 01 10 11 → Bits: 00011011 → Decimal: 27

10 characters × 2 bits = 20 bits total
Fits comfortably in 32-bit integer
```

### Time Complexity Breakdown:
- **Simple Hash**: O(N×10) time, O(N) space - straightforward
- **Rolling Hash**: O(N) time, O(N) space - efficient sliding window
- **Bit Hash**: O(N) time, O(N) space - maximum efficiency
- **Double Hash**: O(N) time, O(N) space - collision reduction

### Alternative Approaches:

#### 1. Bit Manipulation Hash (O(N) time, O(N) space)
```go
func findRepeatedDnaSequencesBitHash(s string) []string {
    // Use 2 bits per character
    // 10 characters = 20 bits, fits in 32-bit int
    mask := (1 << 20) - 1
    
    hash = 0
    for i := 0; i < 10; i++ {
        hash = (hash << 2) | charToBits(s[i])
    }
    
    // Rolling hash with bit operations
    for i := 1; i <= len(s)-10; i++ {
        hash = ((hash << 2) & mask) | charToBits(s[i+9])
        // Check for duplicates...
    }
}
```
- **Pros**: Maximum efficiency, no modulo operations
- **Cons**: Limited to small alphabet, fixed pattern length

#### 2. Double Hashing (O(N) time, O(N) space)
```go
func findRepeatedDnaSequencesDoubleHash(s string) []string {
    // Use two different hash functions
    // Reduces collision probability significantly
    // Store [hash1, hash2] pairs
    // ... implementation details omitted
}
```
- **Pros**: Very low collision probability
- **Cons**: More memory usage, slightly slower

#### 3. Simple String Hash (O(N×10) time, O(N) space)
```go
func findRepeatedDnaSequences(s string) []string {
    // Direct string hashing
    // Store actual strings as keys
    // Simple but less efficient
    // ... implementation details omitted
}
```
- **Pros**: Simple implementation, no collisions
- **Cons**: Slower due to string comparisons

### Extensions for Interviews:
- **Variable Pattern Length**: Generalize to any pattern length
- **Different Alphabets**: Extend to other character sets
- **Counting Variants**: Count occurrences of each repeated sequence
- **Streaming**: Process very large DNA sequences in chunks
- **Real-world Applications**: DNA sequence analysis, plagiarism detection
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Repeated DNA Sequences ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", "Standard case with repeats"},
		{"AAAAAAAAAAAAA", "All same characters"},
		{"ACGTACGTACGTACGTACGTACGT", "No repeats"},
		{"", "Empty string"},
		{"ACGT", "Too short"},
		{"ACGTACGTACGTACGTACGTACGTACGTACGTACGTACGT", "Long string"},
		{"ACGTACGTACACGTACGTAC", "Partial repeats"},
		{"CCCCCCCCCCCCCCCCCCCCCCCC", "All C's"},
		{"AGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCT", "Pattern repeats"},
		{"ACGTACGTACGTACGTACGTACGTACGTACGTACGTACGTACGTACGTACGT", "Very long string"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  String length: %d\n", len(tc.s))
		
		result1 := findRepeatedDnaSequences(tc.s)
		result2 := findRepeatedDnaSequencesRollingHash(tc.s)
		result3 := findRepeatedDnaSequencesBitHash(tc.s)
		result4 := findRepeatedDnaSequencesAlternative(tc.s)
		result5 := findRepeatedDnaSequencesDoubleHash(tc.s)
		result6 := findRepeatedDnaSequencesBruteForce(tc.s)
		
		fmt.Printf("  Simple hash: %v\n", result1)
		fmt.Printf("  Rolling hash: %v\n", result2)
		fmt.Printf("  Bit hash: %v\n", result3)
		fmt.Printf("  Alternative: %v\n", result4)
		fmt.Printf("  Double hash: %v\n", result5)
		fmt.Printf("  Brute force: %v\n\n", result6)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	longString := ""
	for i := 0; i < 10000; i++ {
		longString += "ACGT"
	}
	
	fmt.Printf("Testing with string length: %d\n", len(longString))
	
	// Test different approaches
	result := findRepeatedDnaSequencesRollingHash(longString)
	fmt.Printf("Rolling hash found %d repeated sequences\n", len(result))
	
	result = findRepeatedDnaSequencesBitHash(longString)
	fmt.Printf("Bit hash found %d repeated sequences\n", len(result))
	
	// Test edge cases
	fmt.Println("\n=== Edge Case Testing ===")
	
	// Test with minimum length
	minString := "ACGTACGTAC"
	fmt.Printf("Minimum length string: %v\n", findRepeatedDnaSequences(minString))
	
	// Test with exactly one repeat
	oneRepeat := "ACGTACGTACACGTACGTAC"
	fmt.Printf("One repeat: %v\n", findRepeatedDnaSequences(oneRepeat))
	
	// Test with overlapping repeats
	overlap := "AAAAAAAAAACCCCCCCCCCTTTTTTTTTTGGGGGGGGGGAAAAAAAAAACCCCCCCCCCTTTTTTTTTTGGGGGGGGGG"
	fmt.Printf("Overlapping repeats: %v\n", findRepeatedDnaSequences(overlap))
}
