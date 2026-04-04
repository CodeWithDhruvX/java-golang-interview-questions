package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. Reverse Linked List
// Time: O(N), Space: O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	
	return prev
}

// Helper function to create a linked list from slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Iterative Pointer Reversal
- **Three Pointers**: Previous, Current, Next pointers for reversal
- **In-Place Modification**: Reverse links without creating new nodes
- **Linear Traversal**: Single pass through the list
- **Pointer Manipulation**: Careful Next pointer redirection

## 2. PROBLEM CHARACTERISTICS
- **List Reversal**: Reverse the order of nodes in linked list
- **In-Place Operation**: Modify existing list structure
- **Pointer Redirection**: Change Next pointers to point backward
- **Head Update**: New head is the original tail

## 3. SIMILAR PROBLEMS
- Reverse Linked List II (LeetCode 92) - Reverse sublist
- Reverse Nodes in K-Group (LeetCode 25) - Reverse in groups
- Palindrome Linked List (LeetCode 234) - Check using reversal
- Swap Nodes in Pairs (LeetCode 24) - Pairwise reversal

## 4. KEY OBSERVATIONS
- **Pointer Preservation**: Must save Next pointer before overwriting
- **Prev Tracking**: Need previous node to link backwards
- **Order Independence**: Can reverse from any direction
- **Space Efficiency**: No extra storage needed

## 5. VARIATIONS & EXTENSIONS
- **Recursive Reversal**: Recursive implementation
- **Partial Reversal**: Reverse only part of the list
- **Group Reversal**: Reverse nodes in groups of K
- **Circular Lists**: Handle circular linked lists

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can I modify the list? What about recursion depth?"
- Edge cases: empty list, single node, two nodes
- Time complexity: O(N) time, O(1) space (iterative)
- Recursive alternative: O(N) time, O(N) space (call stack)

## 7. COMMON MISTAKES
- Not saving Next pointer before overwriting
- Losing reference to remaining list
- Not updating prev pointer correctly
- Creating cycles instead of reversal
- Forgetting to return new head (prev)

## 8. OPTIMIZATION STRATEGIES
- **Iterative Approach**: O(N) time, O(1) space
- **Recursive Approach**: O(N) time, O(N) space
- **Tail Recursion**: Optimize recursion (Go doesn't optimize)
- **Early Termination**: Not applicable (need full traversal)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like reversing a train track:**
- You have a series of train cars connected in one direction
- You want to reverse all the connections so they point the opposite way
- You need to carefully disconnect each car and reconnect it backward
- You must keep track of where the remaining cars are while working
- The engine becomes the caboose, and the caboose becomes the engine

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Head of a singly-linked list
2. **Goal**: Reverse the order of nodes
3. **Constraint**: Modify in-place, no new nodes
4. **Output**: New head of reversed list

#### Phase 2: Key Insight Recognition
- **"Three-pointer natural fit"** → Prev, Current, Next
- **"In-place modification"** → Redirect Next pointers
- **"Linear traversal"** → Single pass through list
- **"Head change"** → New head is original tail

#### Phase 3: Strategy Development
```
Human thought process:
"I need to reverse the linked list by changing the Next pointers.
I'll use three pointers: prev (starts nil), current (starts at head), and next.
For each node, I'll save its Next, point it to prev, then move all pointers forward.
When current is nil, prev will be the new head."
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return nil (nothing to reverse)
- **Single node**: Return same node (reversal doesn't change it)
- **Two nodes**: Simple swap of Next pointers
- **Large list**: General algorithm handles all cases

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
List: 1 → 2 → 3 → 4 → 5

Human thinking:
"I'll start with three pointers:
prev = nil, current = 1 → 2 → 3 → 4 → 5, next = nil

Step 1: At node 1
Save next = 1.Next = 2
Point 1.Next = prev = nil
Move prev = current = 1
Move current = next = 2
List: nil ← 1, current = 2 → 3 → 4 → 5

Step 2: At node 2
Save next = 2.Next = 3
Point 2.Next = prev = 1
Move prev = current = 2
Move current = next = 3
List: nil ← 1 ← 2, current = 3 → 4 → 5

Step 3: At node 3
Save next = 3.Next = 4
Point 3.Next = prev = 2
Move prev = current = 3
Move current = next = 4
List: nil ← 1 ← 2 ← 3, current = 4 → 5

Step 4: At node 4
Save next = 4.Next = 5
Point 4.Next = prev = 3
Move prev = current = 4
Move current = next = 5
List: nil ← 1 ← 2 ← 3 ← 4, current = 5

Step 5: At node 5
Save next = 5.Next = nil
Point 5.Next = prev = 4
Move prev = current = 5
Move current = next = nil
List: nil ← 1 ← 2 ← 3 ← 4 ← 5, current = nil

Current is nil, so return prev = 5
Final reversed list: 5 → 4 → 3 → 2 → 1 ✓"
```

#### Phase 6: Intuition Validation
- **Why three pointers work**: Need to save next before overwriting
- **Why prev starts nil**: Original head should point to nil
- **Why O(1) space**: Only three pointers needed
- **Why O(N) time**: Each node processed exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Works but uses O(N) stack space
2. **"Should I create new nodes?"** → No, modify existing list
3. **"What about three pointers?"** → Essential to avoid losing references
4. **"Can I optimize further?"** → Iterative approach is already optimal

### Real-World Analogy
**Like reversing a conga line:**
- You have a line of people holding hands in one direction
- You want to reverse the line so everyone faces the opposite way
- You start at the front and carefully turn each person around
- Each person lets go of the person in front and holds hands with the person behind
- The person at the end becomes the new leader of the line

### Human-Readable Pseudocode
```
function reverseList(head):
    prev = null
    current = head
    
    while current != null:
        next = current.next
        current.next = prev
        prev = current
        current = next
    
    return prev
```

### Execution Visualization

### Example: 1 → 2 → 3 → 4 → 5
```
Pointer Evolution:
Initial: prev=nil, current=1→2→3→4→5, next=nil

Step 1: next=2, 1.next=nil, prev=1, current=2
State: nil←1, current=2→3→4→5

Step 2: next=3, 2.next=1, prev=2, current=3
State: nil←1←2, current=3→4→5

Step 3: next=4, 3.next=2, prev=3, current=4
State: nil←1←2←3, current=4→5

Step 4: next=5, 4.next=3, prev=4, current=5
State: nil←1←2←3←4, current=5

Step 5: next=nil, 5.next=4, prev=5, current=nil
State: nil←1←2←3←4←5, current=nil

Final: return prev=5 (head of reversed list)
```

### Key Visualization Points:
- **Three pointers**: prev, current, next
- **Link reversal**: current.next = prev
- **Pointer movement**: prev = current, current = next
- **Head change**: Return prev as new head

### Memory Layout Visualization:
```
Before reversal:
1 → 2 → 3 → 4 → 5 → nil

During step 3 (reversing node 3):
prev = 2 ← 1 ← nil
current = 3 → 4 → 5
next = 4

After step 3:
prev = 3 ← 2 ← 1 ← nil
current = 4 → 5
next = 4

The reversed portion grows as we process each node.
```

### Time Complexity Breakdown:
- **Iterative Approach**: O(N) time, O(1) space
- **Recursive Approach**: O(N) time, O(N) space (call stack)
- **Array Conversion**: O(N) time, O(N) space (inefficient)
- **Stack-based**: O(N) time, O(N) space (extra stack)

### Alternative Approaches:

#### 1. Recursive Approach (O(N) time, O(N) space)
```go
func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    
    return newHead
}
```
- **Pros**: Elegant, concise code
- **Cons**: O(N) space due to recursion stack

#### 2. Stack-based Approach (O(N) time, O(N) space)
```go
func reverseListStack(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    stack := []*ListNode{}
    current := head
    
    // Push all nodes onto stack
    for current != nil {
        stack = append(stack, current)
        current = current.Next
    }
    
    // Pop and rebuild reversed list
    newHead := stack[len(stack)-1]
    current = newHead
    stack = stack[:len(stack)-1]
    
    for len(stack) > 0 {
        current.Next = stack[len(stack)-1]
        current = current.Next
        stack = stack[:len(stack)-1]
    }
    current.Next = nil
    
    return newHead
}
```
- **Pros**: Conceptually simple
- **Cons**: O(N) space, less efficient

#### 3. Array Conversion (O(N) time, O(N) space)
```go
func reverseListArray(head *ListNode) *ListNode {
    // Convert to array
    values := []int{}
    current := head
    for current != nil {
        values = append(values, current.Val)
        current = current.Next
    }
    
    // Reverse array
    for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
        values[i], values[j] = values[j], values[i]
    }
    
    // Rebuild linked list
    return createLinkedList(values)
}
```
- **Pros**: Simple to understand
- **Cons**: O(N) space, inefficient

### Extensions for Interviews:
- **Reverse Sublist**: Reverse only part of the list
- **Reverse K-Group**: Reverse nodes in groups of K
- **Recursive Variations**: Different recursive approaches
- **Circular Lists**: Handle circular linked list reversal
- **Doubly Linked Lists**: Extend to doubly linked lists
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2},
		{},
		{1},
		{1, 2, 3, 4},
		{5, 4, 3, 2, 1},
	}
	
	for i, nums := range testCases {
		head := createLinkedList(nums)
		reversedHead := reverseList(head)
		result := linkedListToSlice(reversedHead)
		fmt.Printf("Test Case %d: %v -> Reversed: %v\n", i+1, nums, result)
	}
}
