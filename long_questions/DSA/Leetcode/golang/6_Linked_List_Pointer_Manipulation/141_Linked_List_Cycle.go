package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 141. Linked List Cycle
// Time: O(N), Space: O(1) - Floyd's Cycle Detection Algorithm
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	
	slow := head
	fast := head.Next
	
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return false
}

// Helper function to create a linked list with optional cycle
func createLinkedListWithCycle(nums []int, cyclePos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	var cycleNode *ListNode
	
	if cyclePos == 0 {
		cycleNode = head
	}
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
		
		if i == cyclePos {
			cycleNode = current
		}
	}
	
	// Create cycle if cyclePos is valid
	if cyclePos >= 0 && cyclePos < len(nums) {
		current.Next = cycleNode
	}
	
	return head
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Floyd's Cycle Detection (Tortoise and Hare)
- **Two Pointers**: Slow pointer moves 1 step, fast pointer moves 2 steps
- **Cycle Detection**: If pointers meet, cycle exists
- **Space Efficiency**: O(1) space, no extra storage needed
- **Linear Time**: O(N) time complexity

## 2. PROBLEM CHARACTERISTICS
- **Cycle Detection**: Determine if linked list has a cycle
- **Pointer Manipulation**: Navigate through linked list structure
- **Termination Condition**: Meeting of pointers indicates cycle
- **No Modification**: Cannot modify the linked list structure

## 3. SIMILAR PROBLEMS
- Find Start of Cycle (LeetCode 142) - Find where cycle begins
- Happy Number (LeetCode 202) - Similar cycle detection in numbers
- Duplicate Number (LeetCode 287) - Find duplicate using cycle detection
- Circular Linked List - Similar circular structure problems

## 4. KEY OBSERVATIONS
- **Meeting Guarantee**: Fast pointer will eventually catch slow pointer if cycle exists
- **Termination**: Fast pointer reaches end if no cycle
- **Speed Difference**: Fast pointer moves twice as fast as slow pointer
- **Space Optimization**: No hash map needed for visited nodes

## 5. VARIATIONS & EXTENSIONS
- **Find Cycle Start**: Locate where cycle begins
- **Cycle Length**: Calculate length of the cycle
- **Multiple Cycles**: Handle lists with multiple cycles (theoretically impossible)
- **Remove Cycle**: Break the cycle in the linked list

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can I modify the list? What about memory constraints?"
- Edge cases: empty list, single node, no cycle, full cycle
- Time complexity: O(N) time, O(1) space
- Alternative approaches: Hash set with O(N) space

## 7. COMMON MISTAKES
- Not handling empty list or single node cases
- Using wrong pointer initialization (fast should start at head.Next)
- Not checking both fast and fast.Next for nil
- Creating infinite loop if cycle detection logic is wrong
- Using hash map when O(1) space solution exists

## 8. OPTIMIZATION STRATEGIES
- **Floyd's Algorithm**: Optimal O(N) time, O(1) space
- **Hash Set Alternative**: O(N) time, O(N) space (simpler but less optimal)
- **Early Termination**: Fast pointer reaches end quickly if no cycle
- **Pointer Initialization**: Start fast at head.Next for immediate movement

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like two runners on a circular track:**
- You have two runners, one slow and one fast, running on a track
- If the track is straight (no cycle), the fast runner will always finish first
- If the track is circular (has a cycle), the fast runner will eventually lap the slow runner
- When the fast runner laps the slow runner, they meet at the same point
- This meeting proves the track is circular

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Head of a singly-linked list
2. **Goal**: Determine if the list contains a cycle
3. **Constraint**: Cannot modify the list structure
4. **Output**: Boolean indicating cycle existence

#### Phase 2: Key Insight Recognition
- **"Two-pointer natural fit"** → Fast and slow pointers
- **"Meeting indicates cycle"** → Fast catches slow if cycle exists
- **"Space efficiency"** → No extra storage needed
- **"Linear traversal"** → Single pass through list

#### Phase 3: Strategy Development
```
Human thought process:
"I need to detect if there's a cycle in the linked list.
I'll use two pointers: slow moves one step, fast moves two steps.
If there's no cycle, fast will reach the end first.
If there's a cycle, fast will eventually lap slow and they'll meet.
The meeting proves a cycle exists."
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return false (no nodes, no cycle)
- **Single node**: Check if node points to itself
- **No cycle**: Fast pointer reaches end (nil)
- **Full cycle**: Head points back to itself

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
List: 3 → 2 → 0 → -4 → 2 (cycle back to node with value 2)

Human thinking:
"I'll start with slow and fast pointers at head:
slow = 3, fast = 3

Step 1: slow moves to 2, fast moves to 0
slow = 2, fast = 0

Step 2: slow moves to 0, fast moves to 2 (enters cycle)
slow = 0, fast = 2

Step 3: slow moves to -4, fast moves to 0
slow = -4, fast = 0

Step 4: slow moves to 2, fast moves to -4
slow = 2, fast = -4

Step 5: slow moves to 0, fast moves to 2
slow = 0, fast = 2

Step 6: slow moves to -4, fast moves to 0
slow = -4, fast = 0

Eventually, slow and fast will meet at the same node!
This proves there's a cycle ✓"
```

#### Phase 6: Intuition Validation
- **Why two pointers work**: Different speeds guarantee eventual meeting
- **Why O(1) space**: No extra storage needed for visited nodes
- **Why O(N) time**: Fast pointer traverses at most 2N nodes
- **Why meeting indicates cycle**: Only possible if fast laps slow

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use hash set?"** → Works but uses O(N) space
2. **"Should I start fast at head?"** → Better to start at head.Next
3. **"What about pointer speed?"** → Fast should move 2x speed
4. **"Can I optimize further?"** → Floyd's algorithm is already optimal

### Real-World Analogy
**Like detecting if you're walking in circles:**
- You're walking in a maze trying to find the exit
- You leave a trail of bread crumbs to mark where you've been
- If you encounter your own bread crumbs, you know you're walking in circles
- The two-pointer technique is like having two walkers at different speeds
- The faster walker will eventually encounter the slower walker if they're in a loop

### Human-Readable Pseudocode
```
function hasCycle(head):
    if head == null or head.next == null:
        return false
    
    slow = head
    fast = head.next
    
    while fast != null and fast.next != null:
        if slow == fast:
            return true
        slow = slow.next
        fast = fast.next.next
    
    return false
```

### Execution Visualization

### Example: 3 → 2 → 0 → -4 → 2 (cycle)
```
Pointer Movement:
Initial: slow=3, fast=3

Step 1: slow=2, fast=0
Step 2: slow=0, fast=2
Step 3: slow=-4, fast=0
Step 4: slow=2, fast=-4
Step 5: slow=0, fast=2
Step 6: slow=-4, fast=0
...and so on

Eventually slow and fast will meet!
This proves cycle exists.
```

### Key Visualization Points:
- **Speed difference**: Fast moves twice as fast as slow
- **Meeting guarantee**: Fast will catch slow if cycle exists
- **Termination**: Fast reaches end if no cycle
- **Space efficiency**: No extra storage needed

### Memory Layout Visualization:
```
List: 3 → 2 → 0 → -4 → 2 (cycle)
           ↑         ↓
           ←←←←←←←←←←

Pointer positions during traversal:
Step 1: slow=3, fast=3
Step 2: slow=2, fast=0
Step 3: slow=0, fast=2
Step 4: slow=-4, fast=0
Step 5: slow=2, fast=-4
Step 6: slow=0, fast=2

The pointers keep moving around the cycle until they meet.
```

### Time Complexity Breakdown:
- **Floyd's Algorithm**: O(N) time, O(1) space
- **Hash Set Approach**: O(N) time, O(N) space
- **Array Conversion**: O(N) time, O(N) space (inefficient)
- **Recursive Approach**: O(N) time, O(N) space (stack)

### Alternative Approaches:

#### 1. Hash Set Approach (O(N) time, O(N) space)
```go
func hasCycleHashSet(head *ListNode) bool {
    visited := make(map[*ListNode]bool)
    current := head
    
    for current != nil {
        if visited[current] {
            return true
        }
        visited[current] = true
        current = current.Next
    }
    
    return false
}
```
- **Pros**: Simple to understand and implement
- **Cons**: O(N) space, less memory efficient

#### 2. Node Modification (O(N) time, O(1) space) - If allowed
```go
func hasCycleModify(head *ListNode) bool {
    current := head
    
    for current != nil {
        if current.Next == head {
            return true
        }
        next := current.Next
        current.Next = head // Mark as visited
        current = next
    }
    
    return false
}
```
- **Pros**: O(1) space, no extra data structures
- **Cons**: Modifies the list structure (usually not allowed)

#### 3. Recursive Approach (O(N) time, O(N) space)
```go
func hasCycleRecursive(head *ListNode) bool {
    return hasCycleHelper(head, make(map[*ListNode]bool))
}

func hasCycleHelper(node *ListNode, visited map[*ListNode]bool) bool {
    if node == nil {
        return false
    }
    if visited[node] {
        return true
    }
    visited[node] = true
    return hasCycleHelper(node.Next, visited)
}
```
- **Pros**: Elegant recursive solution
- **Cons**: O(N) space due to recursion stack

### Extensions for Interviews:
- **Find Cycle Start**: Locate where cycle begins
- **Cycle Length**: Calculate length of the cycle
- **Remove Cycle**: Break the cycle in the linked list
- **Multiple Lists**: Detect cycles in multiple linked lists
- **Cycle Entry Point**: Find entry point of cycle
*/
func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		cyclePos int // -1 means no cycle
	}{
		{[]int{3, 2, 0, -4}, 1}, // Cycle at position 1 (value 2)
		{[]int{1, 2}, 0},        // Cycle at position 0 (value 1)
		{[]int{1}, -1},          // No cycle
		{[]int{}, -1},           // Empty list
		{[]int{1, 2, 3, 4}, 2},  // Cycle at position 2 (value 3)
		{[]int{1, 2, 3, 4, 5}, -1}, // No cycle
	}
	
	for i, tc := range testCases {
		head := createLinkedListWithCycle(tc.nums, tc.cyclePos)
		hasCycleResult := hasCycle(head)
		
		cycleInfo := "No cycle"
		if tc.cyclePos >= 0 {
			cycleInfo = fmt.Sprintf("Cycle at position %d", tc.cyclePos)
		}
		
		fmt.Printf("Test Case %d: %v (%s) -> Has cycle: %t\n", 
			i+1, tc.nums, cycleInfo, hasCycleResult)
	}
}
