package main

import (
	"fmt"
	"math"
	"strings"
)

// 1. Find longest palindrome substring
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		l := int(math.Max(float64(len1), float64(len2)))
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 2. Find longest common prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

// 3. Check if string contains only digits
func isDigitsOnly(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// 4. Count uppercase & lowercase letters
func countCase(str string) {
	upper, lower := 0, 0
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			upper++
		}
		if char >= 'a' && char <= 'z' {
			lower++
		}
	}
	fmt.Printf("Upper: %d, Lower: %d\n", upper, lower)
}

// 5. Remove special characters
func removeSpecialChars(str string) string {
	var result strings.Builder
	for _, char := range str {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// 6. Find all permutations of string (basic)
func permuteString(str string, l, r int) {
	if l == r {
		fmt.Println(str)
	} else {
		runes := []rune(str)
		for i := l; i <= r; i++ {
			runes[l], runes[i] = runes[i], runes[l]
			permuteString(string(runes), l+1, r)
			runes[l], runes[i] = runes[i], runes[l] // backtrack
		}
	}
}

// 7. Check valid parentheses
func isValidParentheses(s string) bool {
	var stack []rune
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}

// 8. Find duplicate words in string
func findDuplicateWords(sentence string) {
	words := strings.Fields(sentence)
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[strings.ToLower(word)]++
	}
	for word, count := range countMap {
		if count > 1 {
			fmt.Printf("%s ", word)
		}
	}
	fmt.Println()
}

// 9. Count occurrence of given character
func countChar(str string, target rune) int {
	count := 0
	for _, char := range str {
		if char == target {
			count++
		}
	}
	return count
}

// 10. Check if two strings are rotation (Already covered Q19)

// 11. Check pangram
func isPangram(sentence string) bool {
	seen := make(map[rune]bool)
	for _, char := range strings.ToLower(sentence) {
		if char >= 'a' && char <= 'z' {
			seen[char] = true
		}
	}
	return len(seen) == 26
}

// 12. Print all substrings
func printSubstrings(str string) {
	n := len(str)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Println(str[i:j])
		}
	}
}

// 13. Remove consecutive duplicates
func removeConsecutiveDeep(str string) string {
	if len(str) < 2 {
		return str
	}
	var result strings.Builder
	result.WriteByte(str[0])
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			result.WriteByte(str[i])
		}
	}
	return result.String()
}

// 14. Check if strings differ by one character
func differByOne(s1, s2 string) bool {
	if int(math.Abs(float64(len(s1)-len(s2)))) > 1 {
		return false
	}
	countDiff := 0
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			countDiff++
			if countDiff > 1 {
				return false
			}
			if len(s1) > len(s2) {
				i++
			} else if len(s2) > len(s1) {
				j++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}
	if i < len(s1) || j < len(s2) {
		countDiff++
	}
	return countDiff == 1
}

// 15. Find smallest & largest word
func minMaxWords(sentence string) {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return
	}
	minWord, maxWord := words[0], words[0]
	for _, word := range words {
		if len(word) < len(minWord) {
			minWord = word
		}
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	fmt.Printf("Min: %s, Max: %s\n", minWord, maxWord)
}

func main() {
	fmt.Println("1. Longest Palindrome:")
	fmt.Println("babad ->", longestPalindrome("babad"))

	fmt.Println("\n2. Longest Common Prefix:")
	fmt.Println("flower, flow, flight ->", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("\n3. Is Digits Only:")
	fmt.Println("12345 ->", isDigitsOnly("12345"))

	fmt.Println("\n4. Count Case:")
	countCase("Hello World")

	fmt.Println("\n5. Remove Special Chars:")
	fmt.Println("$Hem$lo_World ->", removeSpecialChars("$Hem$lo_World"))

	fmt.Println("\n6. Permutations (ABC):")
	permuteString("ABC", 0, 2)

	fmt.Println("\n7. Valid Parentheses:")
	fmt.Println("((())) ->", isValidParentheses("((()))"))

	fmt.Println("\n8. Duplicate Words:")
	findDuplicateWords("Hello world hello")

	fmt.Println("\n9. Count Char 'l':")
	fmt.Println("hello ->", countChar("hello", 'l'))

	fmt.Println("\n11. Is Pangram:")
	fmt.Println("The quick brown fox jumps over the lazy dog ->", isPangram("The quick brown fox jumps over the lazy dog"))

	fmt.Println("\n12. Substrings (abc):")
	printSubstrings("abc")

	fmt.Println("\n13. Remove Consecutive Dups:")
	fmt.Println("aaabbc ->", removeConsecutiveDeep("aaabbc"))

	fmt.Println("\n14. Differ By One:")
	fmt.Println("abc, abd ->", differByOne("abc", "abd"))

	fmt.Println("\n15. Min Max Words:")
	minMaxWords("This is a test string")
}
