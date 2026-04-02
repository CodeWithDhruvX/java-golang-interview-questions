package main

import (
	"fmt"
	"sort"
)

// 721. Accounts Merge
// Time: O(N log N), Space: O(N)
func accountsMerge(accounts [][]string) [][]string {
	// Map email -> account index
	emailToAccount := make(map[string]int)
	
	// Initialize Union-Find
	n := len(accounts)
	parent := make([]int, n)
	rank := make([]int, n)
	
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	var union func(int, int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return
		}
		
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
	}
	
	// Union accounts with common emails
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if prevAccount, exists := emailToAccount[email]; exists {
				union(i, prevAccount)
			} else {
				emailToAccount[email] = i
			}
		}
	}
	
	// Group emails by root account
	accountEmails := make(map[int][]string)
	for email, accountIdx := range emailToAccount {
		root := find(accountIdx)
		accountEmails[root] = append(accountEmails[root], email)
	}
	
	// Build result
	var result [][]string
	for root, emails := range accountEmails {
		// Sort emails
		sort.Strings(emails)
		
		// Create merged account
		merged := []string{accounts[root][0]}
		merged = append(merged, emails...)
		
		result = append(result, merged)
	}
	
	return result
}

// Alternative approach using DFS
func accountsMergeDFS(accounts [][]string) [][]string {
	// Build adjacency list: email -> account name and connected emails
	emailToName := make(map[string]string)
	emailGraph := make(map[string][]string)
	
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			emailToName[email] = name
			
			// Connect all emails in this account
			for j := i + 1; j < len(account); j++ {
				otherEmail := account[j]
				emailGraph[email] = append(emailGraph[email], otherEmail)
				emailGraph[otherEmail] = append(emailGraph[otherEmail], email)
			}
		}
	}
	
	visited := make(map[string]bool)
	var result [][]string
	
	var dfs func(string, *[]string)
	dfs = func(email string, component *[]string) {
		visited[email] = true
		*component = append(*component, email)
		
		for _, neighbor := range emailGraph[email] {
			if !visited[neighbor] {
				dfs(neighbor, component)
			}
		}
	}
	
	for email := range emailGraph {
		if !visited[email] {
			var component []string
			dfs(email, &component)
			
			// Sort emails
			sort.Strings(component)
			
			// Create merged account
			merged := []string{emailToName[component[0]]}
			merged = append(merged, component...)
			
			result = append(result, merged)
		}
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][][]string{
		{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
		},
		{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin4@m.co", "Kevin2@m.co"},
			{"Ethan", "Ethan5@m.co"},
			{"Hanzo", "Hanzo3@m.co"},
			{"Fern", "Fern5@m.co", "Fern0@m.co"},
		},
		{
			{"Alex", "Alex5@m.co", "Alex1@m.co", "Alex0@m.co"},
			{"Ethan", "Ethan3@m.co", "Ethan4@m.co"},
			{"Kevin", "Kevin4@m.co", "Kevin0@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe1@m.co"},
			{"Gabe", "Gabe3@m.co", "Gabe4@m.co"},
		},
		{
			{"David", "David0@m.co", "David1@m.co"},
			{"David", "David3@m.co", "David4@m.co"},
			{"David", "David4@m.co", "David5@m.co"},
		},
	}
	
	for i, accounts := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Println("Input:")
		for _, account := range accounts {
			fmt.Printf("  %v\n", account)
		}
		
		result1 := accountsMerge(accounts)
		result2 := accountsMergeDFS(accounts)
		
		fmt.Println("\nUnion-Find result:")
		for _, merged := range result1 {
			fmt.Printf("  %v\n", merged)
		}
		
		fmt.Println("\nDFS result:")
		for _, merged := range result2 {
			fmt.Printf("  %v\n", merged)
		}
		fmt.Println()
	}
}
