```go
package main

// 1. Group Anagrams
// Time: O(N * K), Space: O(N * K)
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	
	for _, s := range strs {
		var count [26]int
		for _, char := range s {
			count[char-'a']++
		}
		groups[count] = append(groups[count], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// Brute Force Group Anagrams
// Time: O(N^2 * K), Space: O(N * K)
func GroupAnagramsBruteForce(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	
	var result [][]string
	used := make([]bool, len(strs))
	
	for i := 0; i < len(strs); i++ {
		if !used[i] {
			var group []string
			group = append(group, strs[i])
			used[i] = true
			
			// Find all anagrams of strs[i]
			for j := i + 1; j < len(strs); j++ {
				if !used[j] && areAnagramsBruteForce(strs[i], strs[j]) {
					group = append(group, strs[j])
					used[j] = true
				}
			}
			result = append(result, group)
		}
	}
	return result
}

func areAnagramsBruteForce(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	freq := make(map[rune]int)
	for _, char := range s1 {
		freq[char]++
	}
	for _, char := range s2 {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}
	return true
}
```
