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
```
