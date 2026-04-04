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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Union Find for Email Account Merging
- **Email-to-Account Mapping**: Track which account owns each email
- **Union by Email**: Merge accounts that share common emails
- **Component Grouping**: Group emails by their root account
- **Result Construction**: Build merged accounts with sorted emails

## 2. PROBLEM CHARACTERISTICS
- **Account Structure**: Each account has name and list of emails
- **Email Uniqueness**: Each email belongs to exactly one person
- **Account Merging**: Accounts sharing emails belong to same person
- **Sorted Output**: Emails in each merged account must be sorted

## 3. SIMILAR PROBLEMS
- Friend Circles (LeetCode 547) - Similar connected component detection
- Number of Provinces (LeetCode 547) - Same as Friend Circles
- Redundant Connection (LeetCode 684) - Cycle detection
- Longest Consecutive Sequence (LeetCode 128) - Union Find pattern

## 4. KEY OBSERVATIONS
- **Email as identifier**: Emails uniquely identify account ownership
- **Shared emails indicate same person**: Any shared email means accounts merge
- **Union Find natural fit**: Efficiently track account connectivity
- **Sorting requirement**: Final emails must be in sorted order

## 5. VARIATIONS & EXTENSIONS
- **Case-insensitive emails**: Handle email case variations
- **Account priority**: Choose primary account based on criteria
- **Email validation**: Validate email format
- **Large-scale merging**: Handle millions of accounts efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are emails case-sensitive? Can accounts be empty?"
- Edge cases: empty accounts, single email accounts, circular references
- Time complexity: O(N log N) dominated by sorting
- Space complexity: O(N) for Union Find and email mappings

## 7. COMMON MISTAKES
- Not sorting emails in final result
- Forgetting to handle empty accounts
- Not using email-to-account mapping efficiently
- Missing path compression optimization
- Not handling duplicate emails properly

## 8. OPTIMIZATION STRATEGIES
- **Email-to-account map**: Essential for efficient union operations
- **Path compression**: Critical for performance with many emails
- **Union by rank**: Balances tree height for optimal performance
- **Efficient sorting**: Sort only final email lists, not intermediate

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like merging duplicate user accounts in a system:**
- You have user accounts with names and email addresses
- Same person might have multiple accounts with overlapping emails
- Any shared email means the accounts belong to the same person
- You need to merge all accounts for each person into one
- The merged account should have all unique emails, sorted alphabetically

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: List of accounts, each with name and email list
2. **Goal**: Merge accounts belonging to same person
3. **Output**: List of merged accounts with sorted emails
4. **Constraint**: Accounts sharing any email belong to same person

#### Phase 2: Key Insight Recognition
- **"Email connectivity"** → Emails create connections between accounts
- **"Union Find natural fit"** → Efficiently track account connectivity
- **"Email as key"** → Use emails to identify account relationships
- **"Sorting requirement"** → Final emails must be sorted

#### Phase 3: Strategy Development
```
Human thought process:
"I need to merge accounts that share emails.
I'll use Union Find to track which accounts are connected.
For each email, I'll track which account first owned it.
When I see an email owned by another account, I'll union them.
Finally, I'll group all emails by their root account and sort them."
```

#### Phase 4: Edge Case Handling
- **Empty accounts**: Skip or handle appropriately
- **Single email accounts**: Handle normally
- **Circular references**: Union Find handles automatically
- **Duplicate emails**: Email-to-account map handles duplicates

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Accounts:
["John", "johnsmith@mail.com", "john_newyork@mail.com"]
["John", "johnsmith@mail.com", "john00@mail.com"]  
["Mary", "mary@mail.com"]
["John", "johnnybravo@mail.com"]

Human thinking:
"I'll map emails to account indices:
johnsmith@mail.com → 0
john_newyork@mail.com → 0
johnsmith@mail.com → 1 (already exists!)
john00@mail.com → 1
mary@mail.com → 2
johnnybravo@mail.com → 3

Process unions:
- Account 0 and 1 share johnsmith@mail.com → union(0,1)
- Other accounts have no shared emails

Find root accounts:
- johnsmith@mail.com: find(1) = 0
- john_newyork@mail.com: find(0) = 0  
- john00@mail.com: find(1) = 0
- mary@mail.com: find(2) = 2
- johnnybravo@mail.com: find(3) = 3

Group by root:
Root 0: [johnsmith@mail.com, john_newyork@mail.com, john00@mail.com]
Root 2: [mary@mail.com]
Root 3: [johnnybravo@mail.com]

Build result with sorted emails:
["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"]
["Mary", "mary@mail.com"]
["John", "johnnybravo@mail.com"]"
```

#### Phase 6: Intuition Validation
- **Why Union Find works**: Efficiently tracks dynamic connectivity between accounts
- **Why email mapping works**: Emails uniquely identify account ownership
- **Why O(N log N)**: Union operations are O(N α(N)), sorting dominates
- **Why sorting needed**: Problem requires sorted email output

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DFS?"** → Could, but Union Find is more efficient for this pattern
2. **"Should I sort immediately?"** → No, sort only final merged email lists
3. **"What about duplicate emails?"** → Email-to-account map handles this naturally
4. **"Can I optimize further?"** → Sorting is the bottleneck, Union Find is already optimal

### Real-World Analogy
**Like merging duplicate customer profiles in a CRM system:**
- You have customer profiles with names and email addresses
- Same customer might have multiple profiles with overlapping emails
- Any shared email indicates the profiles belong to the same customer
- You need to merge all profiles for each customer into one master profile
- The merged profile should contain all unique emails, sorted for easy viewing

### Human-Readable Pseudocode
```
function accountsMerge(accounts):
    emailToAccount = map()  // email -> account index
    n = length(accounts)
    
    // Initialize Union-Find
    parent = [0, 1, 2, ..., n-1]
    rank = [0, 0, 0, ..., 0]
    
    function find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])  // Path compression
        return parent[x]
    
    function union(x, y):
        rootX = find(x)
        rootY = find(y)
        if rootX == rootY:
            return
        
        // Union by rank
        if rank[rootX] < rank[rootY]:
            parent[rootX] = rootY
        else if rank[rootX] > rank[rootY]:
            parent[rootY] = rootX
        else:
            parent[rootY] = rootX
            rank[rootX]++
    
    // Union accounts with common emails
    for i, account in enumerate(accounts):
        for email in account[1:]:
            if email in emailToAccount:
                union(i, emailToAccount[email])
            else:
                emailToAccount[email] = i
    
    // Group emails by root account
    accountEmails = map()  // root -> [emails]
    for email, accountIdx in emailToAccount:
        root = find(accountIdx)
        accountEmails[root].append(email)
    
    // Build result
    result = []
    for root, emails in accountEmails:
        sort(emails)  // Sort emails
        merged = [accounts[root][0]]  // Add name
        merged.extend(emails)  // Add sorted emails
        result.append(merged)
    
    return result
```

### Execution Visualization

### Example Accounts:
```
["John", "johnsmith@mail.com", "john_newyork@mail.com"]
["John", "johnsmith@mail.com", "john00@mail.com"]  
["Mary", "mary@mail.com"]
["John", "johnnybravo@mail.com"]
```

### Union Find Process:
```
Email-to-Account Mapping:
johnsmith@mail.com → 0
john_newyork@mail.com → 0
johnsmith@mail.com → 1 (exists!) → union(0,1)
john00@mail.com → 1
mary@mail.com → 2
johnnybravo@mail.com → 3

Union Operations:
union(0,1): find(0)=0, find(1)=1 → parent[1]=0, rank[0]=1

Final Roots:
find(0)=0, find(1)=0, find(2)=2, find(3)=3

Email Grouping:
Root 0: [johnsmith@mail.com, john_newyork@mail.com, john00@mail.com]
Root 2: [mary@mail.com]  
Root 3: [johnnybravo@mail.com]

Sorted Result:
["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"]
["Mary", "mary@mail.com"]
["John", "johnnybravo@mail.com"]
```

### Key Visualization Points:
- **Email mapping**: Track which account first owned each email
- **Union operations**: Merge accounts that share emails
- **Root finding**: Find the representative account for each email
- **Email grouping**: Group all emails by their final root account
- **Sorting**: Sort emails within each merged account

### Memory Layout Visualization:
```
Email-to-Account Map:
{
  "johnsmith@mail.com": 0,
  "john_newyork@mail.com": 0,
  "john00@mail.com": 1,
  "mary@mail.com": 2,
  "johnnybravo@mail.com": 3
}

Parent Array: [0, 0, 2, 3] (after union 0,1)

AccountEmails Grouping:
{
  0: [johnsmith@mail.com, john_newyork@mail.com, john00@mail.com],
  2: [mary@mail.com],
  3: [johnnybravo@mail.com]
}
```

### Time Complexity Breakdown:
- **Email mapping**: O(E) where E is total number of emails
- **Union operations**: O(E α(N)) where α is inverse Ackermann function
- **Root finding**: O(E α(N)) for all emails
- **Email grouping**: O(E) to group by root
- **Sorting**: O(E log E) total for sorting all email lists
- **Total time**: O(E log E) dominated by sorting
- **Space**: O(E + N) for email mapping and Union Find

### Alternative Approaches:

#### 1. DFS Approach (O(N + E + E log E) time, O(N + E) space)
```go
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
```
- **Pros**: More intuitive graph traversal approach
- **Cons**: Building adjacency list can be expensive, recursive stack limitations

#### 2. BFS Approach (O(N + E + E log E) time, O(N + E) space)
```go
func accountsMergeBFS(accounts [][]string) [][]string {
    emailToName := make(map[string]string)
    emailGraph := make(map[string][]string)
    
    // Build graph
    for _, account := range accounts {
        name := account[0]
        for i := 1; i < len(account); i++ {
            email := account[i]
            emailToName[email] = name
            
            for j := i + 1; j < len(account); j++ {
                otherEmail := account[j]
                emailGraph[email] = append(emailGraph[email], otherEmail)
                emailGraph[otherEmail] = append(emailGraph[otherEmail], email)
            }
        }
    }
    
    visited := make(map[string]bool)
    var result [][]string
    
    for email := range emailGraph {
        if !visited[email] {
            queue := []string{email}
            visited[email] = true
            var component []string
            
            for len(queue) > 0 {
                current := queue[0]
                queue = queue[1:]
                component = append(component, current)
                
                for _, neighbor := range emailGraph[current] {
                    if !visited[neighbor] {
                        visited[neighbor] = true
                        queue = append(queue, neighbor)
                    }
                }
            }
            
            sort.Strings(component)
            merged := []string{emailToName[component[0]]}
            merged = append(merged, component...)
            result = append(result, merged)
        }
    }
    
    return result
}
```
- **Pros**: Iterative, no recursion stack
- **Cons**: Similar performance to DFS, queue management overhead

#### 3. Optimized Union Find with Email Indexing (O(N + E log E) time, O(N + E) space)
```go
func accountsMergeOptimized(accounts [][]string) [][]string {
    emailToIndex := make(map[string]int)
    emailToName := make(map[string]string)
    
    n := 0
    for _, account := range accounts {
        name := account[0]
        for i := 1; i < len(account); i++ {
            email := account[i]
            if _, exists := emailToIndex[email]; !exists {
                emailToIndex[email] = n
                emailToName[email] = name
                n++
            }
        }
    }
    
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
    
    // Union emails in same account
    for _, account := range accounts {
        if len(account) <= 1 {
            continue
        }
        
        firstEmail := account[1]
        firstIdx := emailToIndex[firstEmail]
        
        for i := 2; i < len(account); i++ {
            email := account[i]
            emailIdx := emailToIndex[email]
            union(firstIdx, emailIdx)
        }
    }
    
    // Group emails by root
    rootToEmails := make(map[int][]string)
    for email, idx := range emailToIndex {
        root := find(idx)
        rootToEmails[root] = append(rootToEmails[root], email)
    }
    
    // Build result
    var result [][]string
    for _, emails := range rootToEmails {
        sort.Strings(emails)
        merged := []string{emailToName[emails[0]]}
        merged = append(merged, emails...)
        result = append(result, merged)
    }
    
    return result
}
```
- **Pros**: More efficient Union Find with email indexing
- **Cons**: More complex implementation, similar overall complexity

### Extensions for Interviews:
- **Case-insensitive emails**: Normalize email case before processing
- **Account priority**: Choose primary account based on creation date or other criteria
- **Email validation**: Validate email format before processing
- **Large-scale optimization**: Handle millions of accounts efficiently
- **Incremental updates**: Handle adding new accounts efficiently
*/
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
