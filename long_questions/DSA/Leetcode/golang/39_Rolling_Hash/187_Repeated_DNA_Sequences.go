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
