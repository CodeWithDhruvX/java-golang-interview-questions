```go
package main

// 1. Longest Substring Without Repeating Characters
// Time: O(N), Space: O(U)
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	start := 0
	maxLength := 0
	
	for i, char := range s {
		if idx, found := lastSeen[char]; found && idx >= start {
			start = idx + 1
		}
		lastSeen[char] = i
		currentLen := i - start + 1
		if currentLen > maxLength {
			maxLength = currentLen
		}
	}
	return maxLength
}

// 2. Longest Palindromic Substring
// Time: O(N^2), Space: O(1)
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i)
		len2 := expandAroundCenter(i, i+1)
		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}
		
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// 3. Find All Anagrams in a String
// Time: O(N), Space: O(1)
func FindAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}
	
	var pFreq [26]int
	var windowFreq [26]int
	
	for i := 0; i < len(p); i++ {
		pFreq[p[i]-'a']++
		windowFreq[s[i]-'a']++
	}
	
	if pFreq == windowFreq {
		result = append(result, 0)
	}
	
	for i := len(p); i < len(s); i++ {
		windowFreq[s[i]-'a']++
		windowFreq[s[i-len(p)]-'a']--
		
		if pFreq == windowFreq {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

// 4. Smallest Window Containing All Characters
// Time: O(N+M), Space: O(1)
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}
	
	windowFreq := make(map[byte]int)
	left, right := 0, 0
	formed := 0
	required := len(tFreq)
	
	ans := []int{-1, 0, 0} // length, left, right
	
	for right < len(s) {
		char := s[right]
		windowFreq[char]++
		
		if count, exists := tFreq[char]; exists && windowFreq[char] == count {
			formed++
		}
		
		for left <= right && formed == required {
			char = s[left]
			
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}
			
			windowFreq[char]--
			if count, exists := tFreq[char]; exists && windowFreq[char] < count {
				formed--
			}
			left++
		}
		right++
	}
	
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
```
