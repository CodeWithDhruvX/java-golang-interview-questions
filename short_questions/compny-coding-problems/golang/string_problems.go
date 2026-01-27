package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1. Reverse a string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 2. Reverse words in a sentence
func reverseWords(sentence string) string {
	words := strings.Fields(sentence)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// 3. Check if a string is palindrome
func isPalindrome(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// 4. Check if two strings are anagrams
func areAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)
	return strings.Join(s1, "") == strings.Join(s2, "")
}

// 5. Count vowels and consonants
func countVowelsConsonants(str string) {
	vowels := 0
	consonants := 0
	str = strings.ToLower(str)
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			switch char {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				consonants++
			}
		}
	}
	fmt.Printf("Vowels: %d, Consonants: %d\n", vowels, consonants)
}

// 6. Count frequency of characters
func charFrequency(str string) {
	freqMap := make(map[rune]int)
	for _, char := range str {
		freqMap[char]++
	}
	fmt.Println(freqMap)
}

// 7. Find first non-repeating character
func firstNonRepeating(str string) string {
	freqMap := make(map[rune]int)
	for _, char := range str {
		freqMap[char]++
	}
	for _, char := range str {
		if freqMap[char] == 1 {
			return string(char)
		}
	}
	return "NULL"
}

// 8. Remove duplicate characters
func removeDuplicates(str string) string {
	seen := make(map[rune]bool)
	var result strings.Builder
	for _, char := range str {
		if !seen[char] {
			seen[char] = true
			result.WriteRune(char)
		}
	}
	return result.String()
}

// 9. Replace spaces with special character
func replaceSpaces(str, specialChar string) string {
	return strings.ReplaceAll(str, " ", specialChar)
}

// 10. Convert lowercase to uppercase (without built-in)
func toUpperCase(str string) string {
	runes := []rune(str)
	for i, char := range runes {
		if char >= 'a' && char <= 'z' {
			runes[i] = char - 32
		}
	}
	return string(runes)
}

// 11. Find longest word in a string
func longestWord(sentence string) string {
	words := strings.Fields(sentence)
	maxLen := 0
	longest := ""
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
			longest = word
		}
	}
	return longest
}

// 12. Count number of words
func countWords(sentence string) int {
	if len(strings.TrimSpace(sentence)) == 0 {
		return 0
	}
	return len(strings.Fields(sentence))
}

// 13. Check substring present or not
func isSubstring(mainStr, subStr string) bool {
	return strings.Contains(mainStr, subStr)
}

// 14. Remove vowels from string
func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	var result strings.Builder
	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// 15. Sort characters in a string
func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// 16. Find duplicate characters
func findDuplicates(str string) {
	freqMap := make(map[rune]int)
	for _, char := range str {
		freqMap[char]++
	}
	for char, count := range freqMap {
		if count > 1 {
			fmt.Printf("%c ", char)
		}
	}
	fmt.Println()
}

// 17. Reverse string using recursion
func reverseRecursive(str string) string {
	if len(str) == 0 {
		return ""
	}
	return reverseRecursive(str[1:]) + string(str[0])
}

// 18. Print string in zig-zag format
func printZigZag(str string, k int) string {
	if k == 1 {
		return str
	}
	rows := make([]strings.Builder, k)
	row := 0
	down := true
	for _, char := range str {
		rows[row].WriteRune(char)
		if row == 0 {
			down = true
		} else if row == k-1 {
			down = false
		}
		if down {
			row++
		} else {
			row--
		}
	}
	var result strings.Builder
	for _, r := range rows {
		result.WriteString(r.String())
	}
	return result.String()
}

// 19. Check string rotation
func isRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	temp := str1 + str1
	return strings.Contains(temp, str2)
}

// 20. Compare two strings without using equals()
func compareStrings(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("1. Reverse String:")
	fmt.Println("hello ->", reverseString("hello"))

	fmt.Println("\n2. Reverse Words:")
	fmt.Println("Hello World ->", reverseWords("Hello World"))

	fmt.Println("\n3. Is Palindrome:")
	fmt.Println("madam ->", isPalindrome("madam"))
	fmt.Println("hello ->", isPalindrome("hello"))

	fmt.Println("\n4. Are Anagrams:")
	fmt.Println("listen, silent ->", areAnagrams("listen", "silent"))

	fmt.Println("\n5. Count Vowels Consonants:")
	countVowelsConsonants("Hello")

	fmt.Println("\n6. Char Frequency:")
	charFrequency("banana")

	fmt.Println("\n7. First Non-Repeating:")
	fmt.Println("swiss ->", firstNonRepeating("swiss"))

	fmt.Println("\n8. Remove Duplicates:")
	fmt.Println("banana ->", removeDuplicates("banana"))

	fmt.Println("\n9. Replace Spaces:")
	fmt.Println("Hello World, - ->", replaceSpaces("Hello World", "-"))

	fmt.Println("\n10. To Upper Case:")
	fmt.Println("java ->", toUpperCase("java"))

	fmt.Println("\n11. Longest Word:")
	fmt.Println("I love programming ->", longestWord("I love programming"))

	fmt.Println("\n12. Count Words:")
	fmt.Println("Hello World ->", countWords("Hello World"))

	fmt.Println("\n13. Is Substring:")
	fmt.Println("Hello World, World ->", isSubstring("Hello World", "World"))

	fmt.Println("\n14. Remove Vowels:")
	fmt.Println("Hello World ->", removeVowels("Hello World"))

	fmt.Println("\n15. Sort String:")
	fmt.Println("edcba ->", sortString("edcba"))

	fmt.Println("\n16. Find Duplicates:")
	fmt.Print("banana -> ")
	findDuplicates("banana")

	fmt.Println("\n17. Reverse Recursive:")
	fmt.Println("recursion ->", reverseRecursive("recursion"))

	fmt.Println("\n18. Print Zig Zag:")
	fmt.Println("PAYPALISHIRING, 3 ->", printZigZag("PAYPALISHIRING", 3))

	fmt.Println("\n19. Is Rotation:")
	fmt.Println("ABCD, CDAB ->", isRotation("ABCD", "CDAB"))

	fmt.Println("\n20. Compare Strings:")
	fmt.Println("abc, abc ->", compareStrings("abc", "abc"))
}
